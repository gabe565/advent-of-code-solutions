package day2

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
	"github.com/gabe565/advent-of-code-solutions/internal/util"
)

func New() day.Day[[][]int, int] {
	return day.Day[[][]int, int]{
		Year: 2024,
		Day:  2,
		Parse: func(r io.Reader) ([][]int, error) {
			var input [][]int
			scanner := bufio.NewScanner(r)
			for scanner.Scan() {
				line := strings.Fields(scanner.Text())
				result := make([]int, 0, len(line))
				for _, v := range line {
					parsed, err := strconv.Atoi(v)
					if err != nil {
						return nil, err
					}

					result = append(result, parsed)
				}
				input = append(input, result)
			}
			return input, scanner.Err()
		},
		Part1: func(input [][]int) (int, error) {
			return run(input, false)
		},
		Part2: func(input [][]int) (int, error) {
			return run(input, true)
		},
	}
}

func run(input [][]int, hasDampener bool) (int, error) {
	var safeCount int
	for _, row := range input {
		if testRow(row) {
			safeCount++
		} else if hasDampener {
			for i := range row {
				if testRow(slices.Concat(row[0:i], row[i+1:])) {
					safeCount++
					break
				}
			}
		}
	}
	return safeCount, nil
}

func computeDirection(a, b int) int {
	switch {
	case a < b:
		return 1
	case a > b:
		return -1
	default:
		return 0
	}
}

func testRow(row []int) bool {
	direction := computeDirection(row[0], row[1])
	for i, val := range row {
		if i == 0 {
			continue
		}
		prev := row[i-1]
		diff := util.Abs(val - prev)
		if diff < 1 || diff > 3 || (val-prev)*direction < 0 {
			return false
		}
	}
	return true
}
