package utils

import (
	"errors"
	"os"
)

func ParseCommandLineArguments() string {
	args := os.Args[2:]
	if len(args) != 1 {
		panic(errors.New("expected filename"))
	}

	filename := args[0]
	return filename
}
