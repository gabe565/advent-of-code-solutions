package day2

import (
	"bufio"
	"fmt"
	"io"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() *day.Day[[]Game, int] {
	return &day.Day[[]Game, int]{
		Year: 2023,
		Day:  2,
		Parse: func(r io.Reader) ([]Game, error) {
			var games []Game
			scanner := bufio.NewScanner(r)
			for i := 0; scanner.Scan(); i++ {
				var game Game
				if err := game.UnmarshalText(scanner.Bytes()); err != nil {
					return nil, fmt.Errorf("failed to unmarshal game %d: %w", i, err)
				}
				games = append(games, game)
			}
			return games, scanner.Err()
		},
		Part1: func(input []Game) (int, error) {
			var result int
			for _, game := range input {
				if game.Valid() {
					result += game.ID
				}
			}
			return result, nil
		},
		Part2: func(input []Game) (int, error) {
			var result int
			for _, game := range input {
				result += game.Power()
			}
			return result, nil
		},
	}
}
