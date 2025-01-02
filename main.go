package main

import (
	"os"

	"github.com/gabe565/advent-of-code-solutions/cmd"
)

func main() {
	if err := cmd.New().Execute(); err != nil {
		os.Exit(1)
	}
}
