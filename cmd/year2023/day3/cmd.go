package day3

import (
	"io"
	"time"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() *day.Day[*Schematic, int] {
	return &day.Day[*Schematic, int]{
		Date: time.Date(2023, 12, 1, 0, 0, 0, 0, time.Local),
		Parse: func(r io.Reader) (*Schematic, error) {
			var schematic Schematic
			err := schematic.Decode(r)
			return &schematic, err
		},
		Part1: func(input *Schematic) (int, error) {
			result, _ := input.Ratios()
			return result, nil
		},
		Part2: func(input *Schematic) (int, error) {
			_, result := input.Ratios()
			return result, nil
		},
	}
}
