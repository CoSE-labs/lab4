package main

import (
	engine "github.com/CoSE-labs/lab4/engine"
)

func main() {
	EventLoop := new(engine.EventLoop)
	EventLoop.Start()
	cmd := commands.printCommand("TEXT")
	EventLoop.Post(cmd)

	EventLoop.AwaitFinish()
}
