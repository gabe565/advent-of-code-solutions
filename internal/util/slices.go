package util

import "cmp"

func Sum[S ~[]E, E cmp.Ordered](s S) E {
	var result E
	for _, v := range s {
		result += v
	}
	return result
}
