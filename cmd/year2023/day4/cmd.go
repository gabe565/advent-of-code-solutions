package day4

import (
	"io"
	"time"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() *day.Day[Cards, int] {
	return &day.Day[Cards, int]{
		Date: time.Date(2023, 12, 4, 0, 0, 0, 0, time.Local),
		Parse: func(r io.Reader) (Cards, error) {
			var cards Cards
			err := cards.Decode(r)
			return cards, err
		},
		Part1: func(input Cards) (int, error) {
			var result int
			for _, card := range input {
				result += card.Points()
			}
			return result, nil
		},
		Part2: func(input Cards) (int, error) {
			return input.Winning(), nil
		},
	}
}
