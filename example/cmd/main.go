package main

import (
	"fmt"
	"bufio"
	"os"

	"github.com/CoSE-labs/lab4/commands"
	"github.com/CoSE-labs/lab4/engine"
)

func main() {
	file, err := os.Open("./inputFile.txt")
	if err != nil {
		fmt.Println("Unable to open file, ", err)
	}
	defer file.Close()

	eventLoop := new(engine.EventLoop)
	eventLoop.Start()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cmd := commands.Parse(line)
		eventLoop.Post(cmd)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Unable to read file, ", err)
	}
	eventLoop.AwaitFinish()
}
