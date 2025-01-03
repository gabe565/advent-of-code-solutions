package util

import (
	"iter"
	"slices"
)

// Permutations returns a iter.Seq which returns all permutations of a given slice.
//
// Based on https://stackoverflow.com/a/30230552
func Permutations[T any](x []T) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		if len(x) == 0 {
			return
		}
		perm := make([]int, len(x))
		for perm[0] < len(perm) {
			result := slices.Clone(x)
			for i, v := range perm {
				result[i], result[i+v] = result[i+v], result[i]
			}
			if !yield(result) {
				return
			}

			for i := len(perm) - 1; i >= 0; i-- {
				if i == 0 || perm[i] < len(perm)-i-1 {
					perm[i]++
					break
				}
				perm[i] = 0
			}
		}
	}
}
