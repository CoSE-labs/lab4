package commands

import (
	"fmt"

	"github.com/CoSE-labs/lab4/engine"
)

type printCommand struct {
	line string
}

func (p printCommand) Execute(hand engine.Handler) {
	fmt.Println(p.line)
}
