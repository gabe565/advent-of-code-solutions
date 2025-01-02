package util

import (
	"errors"

	"golang.org/x/exp/constraints"
)

func GCD[T constraints.Integer](a, b T) T {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

var ErrLCMArgs = errors.New("LCM requires at least two arguments")

func LCM[T constraints.Integer](n ...T) (T, error) {
	if len(n) < 2 {
		return 0, ErrLCMArgs
	}

	result := n[0] * n[1] / GCD(n[0], n[1])

	for _, v := range n[2:] {
		var err error
		if result, err = LCM(result, v); err != nil {
			return result, err
		}
	}
	return result, nil
}
