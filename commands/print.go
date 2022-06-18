package commands

import (
	"fmt"
)

type PrintCommand struct {
	line string
}

func (p PrintCommand) Execute(hand Handler) {
	fmt.Println(p.line)
}
