package day1

import (
	"bufio"
	"io"
	"strconv"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

const WindowWidth = 3

func New() *day.Day[[]int, int] {
	return &day.Day[[]int, int]{
		Year: 2021,
		Day:  1,
		Parse: func(r io.Reader) ([]int, error) {
			scanner := bufio.NewScanner(r)
			var vals []int
			for scanner.Scan() {
				val, err := strconv.Atoi(scanner.Text())
				if err != nil {
					return nil, err
				}

				vals = append(vals, val)
			}
			return vals, scanner.Err()
		},
		Part1: func(input []int) (int, error) {
			var count int
			prev := -1
			for _, v := range input {
				switch {
				case prev == -1:
				case prev < v:
					count++
				}
				prev = v
			}

			return count, nil
		},
		Part2: func(input []int) (int, error) {
			var count int
			prev := -1
		nums:
			for i, num := range input {
				curr := num
				for j := i + 1; j < i+WindowWidth; j++ {
					if j < len(input) {
						curr += input[j]
					} else {
						break nums
					}
				}

				switch {
				case prev == -1:
				case prev < curr:
					count++
				}
				prev = curr
			}

			return count, nil
		},
	}
}
