package day2

import (
	"io"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() *day.Day[io.Reader, int] {
	return &day.Day[io.Reader, int]{
		Year: 2021,
		Day:  2,
		Parse: func(r io.Reader) (io.Reader, error) {
			return r, nil
		},
		Part1: func(input io.Reader) (int, error) {
			var point Point
			if _, err := io.Copy(&point, input); err != nil {
				return 0, err
			}

			return point.Multiply(), nil
		},
		Part2: func(input io.Reader) (int, error) {
			point := Point{Mode: MoveAim}
			if _, err := io.Copy(&point, input); err != nil {
				return 0, err
			}
			return point.Multiply(), nil
		},
	}
}
