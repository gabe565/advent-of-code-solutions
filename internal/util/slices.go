package util

import (
	"cmp"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

func Sum[S ~[]E, E cmp.Ordered](s S) E {
	var result E
	for _, v := range s {
		result += v
	}
	return result
}

func Multiply[S ~[]E, E constraints.Integer | constraints.Float](s S) E {
	result := E(1)
	for _, v := range s {
		result *= v
	}
	return result
}

func StringToIntSlice(s string, sep string) ([]int, error) {
	split := strings.Split(s, sep)
	result := make([]int, 0, len(split))
	for _, v := range split {
		if v != "" {
			v, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			result = append(result, v)
		}
	}
	return result, nil
}

func Count[T comparable](s []T, search T) int {
	var count int
	for _, v := range s {
		if v == search {
			count++
		}
	}
	return count
}
