package day3

import (
	"io"
	"regexp"
	"strconv"
	"time"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() day.Day[[]byte, int] {
	return day.Day[[]byte, int]{
		Date:  time.Date(2024, 12, 3, 0, 0, 0, 0, time.Local),
		Parse: io.ReadAll,
		Part1: func(input []byte) (int, error) {
			re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
			var result int
			for _, match := range re.FindAllSubmatch(input, -1) {
				num, err := strconv.Atoi(string(match[1]))
				if err != nil {
					return 0, err
				}

				multiplier, err := strconv.Atoi(string(match[2]))
				if err != nil {
					return 0, err
				}

				result += num * multiplier
			}
			return result, nil
		},
		Part2: func(input []byte) (int, error) {
			re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do(?:n't)?\(\)`)
			var result int
			enabled := true
			for _, match := range re.FindAllSubmatch(input, -1) {
				switch string(match[0]) {
				case "do()":
					enabled = true
				case "don't()":
					enabled = false
				default:
					if enabled {
						num, err := strconv.Atoi(string(match[1]))
						if err != nil {
							return 0, err
						}

						multiplier, err := strconv.Atoi(string(match[2]))
						if err != nil {
							return 0, err
						}

						result += num * multiplier
					}
				}
			}
			return result, nil
		},
	}
}
