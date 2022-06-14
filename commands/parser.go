package commands

import (
	"strings"

	"github.com/CoSE-labs/lab4/engine"
)

func Parse(fileline string) engine.Command {
	getParts := strings.Fields(fileline)

	if len(getParts) != 2 {
		return &printCommand{line: "Syntax Error: invalid number of arguments"}
	}

	com := getParts[0]
	argument := getParts[1]

	switch com {
	case "print":
		return &printCommand{line: argument}
	case "reverse":
		return &revCommand{arg: argument}
	default:
		return &printCommand{line: "Syntax Error: unknown command"}
	}
}
