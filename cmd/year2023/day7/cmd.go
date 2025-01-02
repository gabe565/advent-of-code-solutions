package day7

import (
	"io"
	"time"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() *day.Day[*Game, int] {
	return &day.Day[*Game, int]{
		Date: time.Date(2023, 12, 7, 0, 0, 0, 0, time.Local),
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
