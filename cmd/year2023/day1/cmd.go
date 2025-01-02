package day1

import (
	"bytes"
	"io"
	"strings"
	"time"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() *day.Day[[]string, int] {
	return &day.Day[[]string, int]{
		Date: time.Date(2023, 12, 1, 0, 0, 0, 0, time.Local),
		Parse: func(r io.Reader) ([]string, error) {
			b, err := io.ReadAll(r)
			if err != nil {
				return nil, err
			}
			b = bytes.TrimSpace(b)
			return strings.Split(string(b), "\n"), nil
		},
		Part1: func(input []string) (int, error) {
			var result int
			for _, line := range input {
				result += findFirstLast(line)
			}
			return result, nil
		},
		Part2: func(input []string) (int, error) {
			replacer := strings.NewReplacer(
				"one", "o1e",
				"two", "t2o",
				"three", "t3e",
				"four", "f4r",
				"five", "f5e",
				"six", "s6x",
				"seven", "s7n",
				"eight", "e8t",
				"nine", "n9e",
			)

			var result int
			for _, line := range input {
				for {
					replaced := replacer.Replace(line)
					if line == replaced {
						break
					}
					line = replaced
				}
				result += findFirstLast(line)
			}
			return result, nil
		},
	}
}

func findFirstLast(line string) int {
	var first *int
	var second int
	for _, b := range []byte(line) {
		if b >= '0' && b <= '9' {
			i := int(b - '0')
			if first == nil {
				first = &i
			}
			second = i
		}
	}

	if first == nil {
		return 0
	}
	return 10*(*first) + second
}
