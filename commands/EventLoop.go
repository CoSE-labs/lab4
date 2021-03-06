package commands

import "sync"

type commandsQueue struct {
	sync.Mutex

	CmdArray []Command
	wait     bool
	notEmpty chan struct{}
}

type cmdExecutor struct {
	executor func()
}

type EventLoop struct {
	queue      *commandsQueue
	stop       bool
	stopSignal chan struct{}
}

func (cq *commandsQueue) push(cmd Command) {
	cq.Lock()
	defer cq.Unlock()

	cq.CmdArray = append(cq.CmdArray, cmd)

	if cq.wait {
		cq.wait = false
		cq.notEmpty <- struct{}{}
	}
}

func (cq *commandsQueue) pull() Command {
	cq.Lock()
	defer cq.Unlock()

	if len(cq.CmdArray) == 0 {
		cq.wait = true
		cq.Unlock()
		<-cq.notEmpty
		cq.Lock()
	}

	res := cq.CmdArray[0]
	cq.CmdArray[0] = nil
	cq.CmdArray = cq.CmdArray[1:]

	return res
}

func (cq *commandsQueue) empty() int {
	cq.Lock()
	defer cq.Unlock()

	return len(cq.CmdArray)
}

func (ce *cmdExecutor) Execute(h Handler) {
	ce.executor()
}

func (loop *EventLoop) Start() {
	loop.queue = &commandsQueue{
		notEmpty: make(chan struct{}),
	}
	loop.stopSignal = make(chan struct{})

	go func() {
		for !(loop.stop && loop.queue.empty() == 0) {
			cmd := loop.queue.pull()
			cmd.Execute(loop)
		}
		loop.stopSignal <- struct{}{}
	}()
}

func (loop *EventLoop) Post(cmd Command) {
	loop.queue.push(cmd)
}

func (loop *EventLoop) AwaitFinish() {
	finish := &cmdExecutor{func() { loop.stop = true }}
	loop.Post(finish)

	<-loop.stopSignal
}
