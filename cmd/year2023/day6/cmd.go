package day6

import (
	"io"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() *day.Day[*Race, int] {
	return &day.Day[*Race, int]{
		Year: 2023,
		Day:  6,
		Parse: func(r io.Reader) (*Race, error) {
			var race Race
			err := race.Decode(r)
			return &race, err
		},
		Part1: func(input *Race) (int, error) {
			return input.Part1()
		},
		Part2: func(input *Race) (int, error) {
			return input.Part2()
		},
	}
}
