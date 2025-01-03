package day1

import (
	"bufio"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
	"github.com/gabe565/advent-of-code-solutions/internal/util"
	"golang.org/x/exp/slices"
)

func New() day.Day[[][]int, int] {
	return day.Day[[][]int, int]{
		Date: time.Date(2024, 12, 1, 0, 0, 0, 0, time.Local),
		Parse: func(r io.Reader) ([][]int, error) {
			input := make([][]int, 2)
			scanner := bufio.NewScanner(r)
			for scanner.Scan() {
				line := strings.Fields(scanner.Text())
				for i, v := range line {
					parsed, err := strconv.Atoi(v)
					if err != nil {
						return nil, err
					}

					input[i] = append(input[i], parsed)
				}
			}
			return input, scanner.Err()
		},
		Part1: func(input [][]int) (int, error) {
			input = [][]int{
				slices.Clone(input[0]),
				slices.Clone(input[1]),
			}
			slices.Sort(input[0])
			slices.Sort(input[1])

			var result int
			for i := range input[0] {
				result += util.Abs(input[0][i] - input[1][i])
			}
			return result, nil
		},
		Part2: func(input [][]int) (int, error) {
			var result int
			for _, needle := range input[0] {
				var count int
				for _, v := range input[1] {
					if v == needle {
						count++
					}
				}
				result += count * needle
			}
			return result, nil
		},
	}
}
