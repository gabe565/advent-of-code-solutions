package day6

import (
	"fmt"
	"io"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
	"github.com/gabe565/advent-of-code-solutions/internal/util"
)

func New() *day.Day[[]int, int] {
	return &day.Day[[]int, int]{
		Year: 2021,
		Day:  6,
		Parse: func(r io.Reader) ([]int, error) {
			data, err := io.ReadAll(r)
			if err != nil {
				return nil, err
			}

			input, err := util.StringToIntSlice(string(data), ",")
			if err != nil {
				return nil, err
			}

			for _, n := range input {
				if n < 0 || n > 8 {
					return nil, fmt.Errorf("invalid input: %d", n)
				}
			}

			return input, nil
		},
		Part1: func(input []int) (int, error) {
			return run(input, 80), nil
		},
		Part2: func(input []int) (int, error) {
			return run(input, 256), nil
		},
	}
}

func run(input []int, days int) int {
	var counts [9]int
	for _, n := range input {
		counts[n]++
	}

	for range days {
		zeros := counts[0]
		copy(counts[:], counts[1:])
		counts[6] += zeros
		counts[8] = zeros
	}

	return util.Sum(counts[:])
}
