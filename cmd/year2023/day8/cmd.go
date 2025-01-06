package day8

import (
	"io"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() *day.Day[*Network, int] {
	return &day.Day[*Network, int]{
		Year: 2023,
		Day:  8,
		Parse: func(r io.Reader) (*Network, error) {
			var network Network
			err := network.Decode(r)
			return &network, err
		},
		Part1: func(input *Network) (int, error) {
			return input.Steps("AAA", false)
		},
		Part2: func(input *Network) (int, error) {
			return input.GhostSteps()
		},
	}
}
