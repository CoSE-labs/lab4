package engine

import "sync"

type commandsQueue struct {
	mu       sync.Mutex
	cmdArray []Command
	wait     bool

	notEmpty chan struct{}
}

func (cq *commandsQueue) push(cmd Command) {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	cq.cmdArray = append(cq.cmdArray, cmd)
	if cq.wait {
		cq.notEmpty <- struct{}{}
	}
}

func (cq *commandsQueue) pull() Command {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	if len(cq.cmdArray == 0) {
		cq.wait = true
		cq.mu.Unlock()
		<-cq.notEmpty
		cq.mu.Lock()
	}
	res := cq.cmdArray[0]
	cq.cmdArray[0] = nil
	cq.cmdArray = cq.cmdArray[1:]
	return res
}

func (cq *commandsQueue) empty() bool {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	return len(cq.cmdArray) == 0
}

type EventLoop struct {
	queue *commandsQueue

	stop       bool
	stopSignal chan struct{}
}

func (loop *EventLoop) Start() {
	loop.queue = &commandsQueue{
		notEmpty: make(chan struct{}),
	}
	loop.stopSignal = make(chan struct{})
	go func() {
		for loop.stop != true && loop.queue.empty() {
			cmd := loop.queue.pull()
			cmd.Execute(loop)
		}
		loop.stopSignal <- struct{}{}

	}()
}

func (loop *EventLoop) Post(cmd Command) {
	loop.queue.push(cmd)
}

type CommandFunc func(h Handler)

func (cf CommandFunc) Execute(hand Handler) {
	cf(hand)
}

func (loop *EventLoop) AwaitFinish() {
	loop.Post(CommandFunc(func(h Handler) {
		loop.stop = true
	}))
	<-loop.stopSignal
}
