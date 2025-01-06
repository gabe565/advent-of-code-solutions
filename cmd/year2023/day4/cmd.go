package day4

import (
	"io"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() *day.Day[Cards, int] {
	return &day.Day[Cards, int]{
		Year: 2023,
		Day:  4,
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
