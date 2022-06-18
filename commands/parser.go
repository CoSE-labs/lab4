package commands

import (
	"strings"
)

func Parse(fileline string) Command {
	getParts := strings.Fields(fileline)

	if len(getParts) != 2 {
		return &PrintCommand{line: "Syntax Error: invalid number of arguments"}
	}

	com := getParts[0]
	argument := getParts[1]

	switch com {
	case "print":
		return &PrintCommand{line: argument}
	case "reverse":
		return &RevCommand{arg: argument}
	default:
		return &PrintCommand{line: "Syntax Error: unknown command"}
	}
}
