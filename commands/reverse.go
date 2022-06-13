package commands

import (
	"github.com/CoSE-labs/lab4/engine"
)

type revCommand string

func (revC revCommand) Execute(loop engine.Handler) {
	var result string
	for _, v := range revC {
		result = string(v) + result
	}
	loop.Post(&printCommand(result))
}
