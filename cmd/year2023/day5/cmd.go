package day5

import (
	"io"
	"slices"
	"time"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() *day.Day[*Almanac, int] {
	return &day.Day[*Almanac, int]{
		Date: time.Date(2023, 12, 5, 0, 0, 0, 0, time.Local),
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
