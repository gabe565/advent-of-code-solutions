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

func Abs[T constraints.Signed | constraints.Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Pow returns x**y, the base-x exponential of y.
//
// Based on http://www.programminglogic.com/fast-exponentiation-algorithms/
func Pow[T constraints.Integer](x, y T) T {
	result := T(1)
	for y != 0 {
		if y&1 == 1 {
			result *= x
		}
		y >>= 1
		x *= x
	}
	return result
}
