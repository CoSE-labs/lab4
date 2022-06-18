package commands

type RevCommand struct {
	arg string
}

func (revC *RevCommand) Execute(loop Handler) {
	var result string
	for _, v := range revC.arg {
		result = string(v) + result
	}
	loop.Post(&PrintCommand{line: result})
}
