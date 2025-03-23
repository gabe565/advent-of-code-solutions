package day4

import "fmt"

type Cell struct {
	Value int
	Drawn bool
}

func (c Cell) String() string {
	var result string
	if c.Drawn {
		result += "\033[1m"
	}
	result += fmt.Sprintf("%2d", c.Value) + " "
	if c.Drawn {
		result += "\033[0m"
	}
	return result
}
