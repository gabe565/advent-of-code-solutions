package day11

import (
	"bytes"
	"io"
	"math"
	"strconv"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() day.Day[Stones, int] {
	return day.Day[Stones, int]{
		Year: 2024,
		Day:  11,
		Parse: func(r io.Reader) (Stones, error) {
			b, err := io.ReadAll(r)
			if err != nil {
				return nil, err
			}

			fields := bytes.Fields(bytes.TrimSpace(b))
			stones := make(Stones, 0, len(fields))
			for _, field := range fields {
				parsed, err := strconv.ParseInt(string(field), 10, 64)
				if err != nil {
					return nil, err
				}
				stones = append(stones, Stone(parsed))
			}
			return stones, nil
		},
		Part1: func(input Stones) (int, error) {
			return input.FindCount(25), nil
		},
		Part2: func(input Stones) (int, error) {
			return input.FindCount(75), nil
		},
	}
}

type Stone int64

type Stones []Stone

func (stones Stones) FindCount(n int) int {
	var result int
	for _, b := range stones {
		result += b.FindCount(n)
	}
	return result
}

//nolint:gochecknoglobals
var memoizer = make(map[memoizerKey]int)

type memoizerKey struct {
	s Stone
	n int
}

func (s Stone) FindCount(n int) int {
	if n == 0 {
		return 1
	} else if result, ok := memoizer[memoizerKey{s, n}]; ok {
		return result
	}

	var result int
	switch {
	case s == 0:
		result = (s + 1).FindCount(n - 1)
	case len(strconv.FormatInt(int64(s), 10))%2 == 0:
		digits := int(math.Log10(float64(s)) + 1)
		left := int64(s) / int64(math.Pow10(digits/2))
		right := int64(s) % int64(math.Pow10(digits/2))
		result = Stone(left).FindCount(n-1) + Stone(right).FindCount(n-1)
	default:
		result = (s * 2024).FindCount(n - 1)
	}
	memoizer[memoizerKey{s, n}] = result
	return result
}
