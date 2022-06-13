package main

import (
	"bufio"
	engine "github.com/CoSE-labs/lab4/engine"
	"os"
)

func main() {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	if input, err := os.Open(inputFile); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := parse(commandLine) // parse the line to get a Command
			eventLoop.Post(cmd)
		}
	}
	eventLoop.AwaitFinish()
}
