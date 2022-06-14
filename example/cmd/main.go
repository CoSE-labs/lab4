package main

import (
	"bufio"
	"log"
	"os"

	"github.com/CoSE-labs/lab4/commands"
	"github.com/CoSE-labs/lab4/engine"
)

func main() {
	file, err := os.Open("./inputFile.txt")
	if err != nil {
		log.Fatal("Unable to open file, %s", err)
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
		log.Fatal("Unable to read file, %s", err)
	}
	eventLoop.AwaitFinish()
}
