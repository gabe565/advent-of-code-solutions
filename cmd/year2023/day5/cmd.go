package day5

import (
	"io"
	"slices"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() *day.Day[*Almanac, int] {
	return &day.Day[*Almanac, int]{
		Year: 2023,
		Day:  5,
		Parse: func(r io.Reader) (*Almanac, error) {
			var almanac Almanac
			err := almanac.Decode(r)
			return &almanac, err
		},
		Part1: func(input *Almanac) (int, error) {
			return slices.Min(input.Locations()), nil
		},
		Part2: func(input *Almanac) (int, error) {
			return input.MinLocationRange(), nil
		},
	}
}
