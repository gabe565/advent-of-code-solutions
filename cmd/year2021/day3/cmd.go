package day3

import (
	"io"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() *day.Day[io.Reader, int] {
	return &day.Day[io.Reader, int]{
		Year: 2021,
		Day:  3,
		Parse: func(r io.Reader) (io.Reader, error) {
			return r, nil
		},
		Part1: func(input io.Reader) (int, error) {
			d := NewDiagnostic()
			if _, err := io.Copy(d, input); err != nil {
				return 0, err
			}

			return d.PowerConsumption(), nil
		},
		Part2: func(input io.Reader) (int, error) {
			d := NewLifeSupportDiagnostic()
			if _, err := io.Copy(d, input); err != nil {
				return 0, err
			}

			return d.LifeSupport()
		},
	}
}
