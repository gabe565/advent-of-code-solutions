package day8

import (
	"io"
	"time"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() *day.Day[*Network, int] {
	return &day.Day[*Network, int]{
		Date: time.Date(2023, 12, 8, 0, 0, 0, 0, time.Local),
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
