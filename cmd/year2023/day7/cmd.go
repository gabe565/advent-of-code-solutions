package day7

import (
	"io"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() *day.Day[*Game, int] {
	return &day.Day[*Game, int]{
		Year: 2023,
		Day:  7,
		Parse: func(r io.Reader) (*Game, error) {
			var game Game
			err := game.Decode(r)
			return &game, err
		},
		Part1: func(input *Game) (int, error) {
			return input.Winnings(false), nil
		},
		Part2: func(input *Game) (int, error) {
			return input.Winnings(true), nil
		},
	}
}
