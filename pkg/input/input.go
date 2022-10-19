package input

import "os"

type InputArguments struct {
	Filename string
}

func ProcessArguments() InputArguments {
	args := os.Args[1:]

	if len(args) == 0 {
		panic("No arguments provided")
	}

	arguments := InputArguments{
		Filename: args[0],
	}

	return arguments
}
