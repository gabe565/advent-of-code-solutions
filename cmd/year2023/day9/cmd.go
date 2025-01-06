package day9

import (
	"io"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() *day.Day[*Report, int] {
	return &day.Day[*Report, int]{
		Year: 2023,
		Day:  9,
		Parse: func(r io.Reader) (*Report, error) {
			var report Report
			err := report.Decode(r)
			return &report, err
		},
		Part1: func(input *Report) (int, error) {
			return input.Predict(PredictFuture), nil
		},
		Part2: func(input *Report) (int, error) {
			return input.Predict(PredictPast), nil
		},
	}
}
