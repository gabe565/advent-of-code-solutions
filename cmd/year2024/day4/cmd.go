package day4

import (
	"bytes"
	"io"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() day.Day[[][]byte, int] {
	return day.Day[[][]byte, int]{
		Year: 2024,
		Day:  4,
		Parse: func(r io.Reader) ([][]byte, error) {
			b, err := io.ReadAll(r)
			if err != nil {
				return nil, err
			}
			b = bytes.TrimSpace(b)
			return bytes.Split(b, []byte("\n")), nil
		},
		Part1: func(input [][]byte) (int, error) {
			var result int
			for y, line := range input {
				for x, b := range line {
					if b == 'X' {
						if checkNorth(input, x, y) {
							result++
						}
						if checkNorthEast(input, x, y) {
							result++
						}
						if checkEast(input, x, y) {
							result++
						}
						if checkSouthEast(input, x, y) {
							result++
						}
						if checkSouth(input, x, y) {
							result++
						}
						if checkSouthWest(input, x, y) {
							result++
						}
						if checkWest(input, x, y) {
							result++
						}
						if checkNorthWest(input, x, y) {
							result++
						}
					}
				}
			}
			return result, nil
		},
		Part2: func(input [][]byte) (int, error) {
			var result int
			for y, line := range input {
				for x, b := range line {
					if b == 'A' && checkX(input, x, y) {
						result++
					}
				}
			}
			return result, nil
		},
	}
}

func checkNorth(input [][]byte, x, y int) bool {
	return y >= 3 &&
		input[y-1][x] == 'M' && input[y-2][x] == 'A' && input[y-3][x] == 'S'
}

func checkNorthEast(input [][]byte, x, y int) bool {
	return y >= 3 && x+3 < len(input[y]) &&
		input[y-1][x+1] == 'M' && input[y-2][x+2] == 'A' && input[y-3][x+3] == 'S'
}

func checkEast(input [][]byte, x, y int) bool {
	return x+3 < len(input[y]) &&
		input[y][x+1] == 'M' && input[y][x+2] == 'A' && input[y][x+3] == 'S'
}

func checkSouthEast(input [][]byte, x, y int) bool {
	return y+3 < len(input) && x+3 < len(input[y]) &&
		input[y+1][x+1] == 'M' && input[y+2][x+2] == 'A' && input[y+3][x+3] == 'S'
}

func checkSouth(input [][]byte, x, y int) bool {
	return y+3 < len(input) &&
		input[y+1][x] == 'M' && input[y+2][x] == 'A' && input[y+3][x] == 'S'
}

func checkWest(input [][]byte, x, y int) bool {
	return x >= 3 &&
		input[y][x-1] == 'M' && input[y][x-2] == 'A' && input[y][x-3] == 'S'
}

func checkSouthWest(input [][]byte, x, y int) bool {
	return x >= 3 && y+3 < len(input) &&
		input[y+1][x-1] == 'M' && input[y+2][x-2] == 'A' && input[y+3][x-3] == 'S'
}

func checkNorthWest(input [][]byte, x, y int) bool {
	return y >= 3 && x >= 3 &&
		input[y-1][x-1] == 'M' && input[y-2][x-2] == 'A' && input[y-3][x-3] == 'S'
}

func checkX(input [][]byte, x, y int) bool {
	return y > 0 && x > 0 && y+1 < len(input) && x+1 < len(input[y]) && // Bounds check
		(input[y-1][x-1] == 'M' && input[y+1][x+1] == 'S' || input[y+1][x+1] == 'M' && input[y-1][x-1] == 'S') && // NW, SE
		(input[y-1][x+1] == 'M' && input[y+1][x-1] == 'S' || input[y+1][x-1] == 'M' && input[y-1][x+1] == 'S') // NE, SW
}
