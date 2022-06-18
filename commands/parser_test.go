package commands

import (
	"fmt"
	"testing"
)

var testCases = []string{
	"print 0Hello!",
	"print print",
	"printc calculate",
	"print print print",
	"reverse Github!",
	"reverse reverse reverse",
}

var wantOutputs = []string{
	"0Hello!",
	"print",
	"Syntax Error: unknown command",
	"Syntax Error: invalid number of arguments",
	"!buhtiG",
	"Syntax Error: invalid number of arguments",
}

func TestParse(t *testing.T) {
	var want []interface{}
	var got []interface{}

	for _, value := range testCases {
		outputValue := Parse(value)
		got = append(got, outputValue)
		fmt.Println(outputValue)
	}

	for _, value := range testCases {
		outputValue := Parse(value)
		got = append(got, outputValue)
		fmt.Println(Parse(value))
	}

	for i := range want {
		if want[i] == got[i] {
			continue
		} else {
			t.Errorf("Failed testing Parse(): got %q, want %q", got[i], want[i])
		}
	}
}
