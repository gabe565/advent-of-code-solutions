package day4

import (
	"io"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() *day.Day[*Bingo, int] {
	return &day.Day[*Bingo, int]{
		Year: 2021,
		Day:  4,
		Parse: func(r io.Reader) (*Bingo, error) {
			bingo := &Bingo{}
			_, err := io.Copy(bingo, r)
			return bingo, err
		},
		Part1: func(input *Bingo) (int, error) {
			var board *Board
			for input.Draw() {
				_, board = input.Winner()
				if board != nil {
					break
				}
			}

			sum := board.Sum()
			drawn := input.Order[input.DrawIndex-1]
			return sum * drawn, nil
		},
		Part2: func(input *Bingo) (int, error) {
			winOrder := make([]int, len(input.Boards))
			var wins int
			for input.Draw() {
				done := true
				for i, board := range input.Boards {
					if winOrder[i] == 0 {
						if board.Wins() {
							wins++
							winOrder[i] = wins
						} else {
							done = false
						}
					}
				}
				if done {
					break
				}
			}

			var maxVal int
			var idx int
			for k, v := range winOrder {
				if v > maxVal {
					idx = k
					maxVal = v
				}
			}

			sum := input.Boards[idx].Sum()
			drawn := input.Order[input.DrawIndex-1]
			return sum * drawn, nil
		},
	}
}
