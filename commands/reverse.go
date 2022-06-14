package commands

import (
	"github.com/CoSE-labs/lab4/engine"
)

type revCommand struct {
	arg string
}

func (revC *revCommand) Execute(loop engine.Handler) {
	var result string
	for _, v := range revC.arg {
		result = string(v) + result
	}
	loop.Post(&printCommand{line: result})
}
