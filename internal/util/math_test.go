package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/constraints"
)

func TestGCD(t *testing.T) {
	type args[T constraints.Integer] struct {
		a T
		b T
	}
	type testCase[T constraints.Integer] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{"2 4", args[int]{2, 4}, 2},
		{"3 4", args[int]{3, 4}, 1},
		{"30 40", args[int]{30, 40}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, GCD(tt.args.a, tt.args.b))
		})
	}
}

func TestLCM(t *testing.T) {
	type args[T constraints.Integer] struct {
		n []T
	}
	type testCase[T constraints.Integer] struct {
		name    string
		args    args[T]
		want    T
		wantErr require.ErrorAssertionFunc
	}
	tests := []testCase[int]{
		{"3 4", args[int]{[]int{3, 4}}, 12, require.NoError},
		{"8 16 80", args[int]{[]int{2, 4, 8}}, 8, require.NoError},
		{"1 3 5", args[int]{[]int{1, 3, 5}}, 15, require.NoError},
		{"12 15 75", args[int]{[]int{12, 15, 75}}, 300, require.NoError},
		{"1 2 3 4 5 6 7 8 9", args[int]{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}}, 2520, require.NoError},
		{"one arg", args[int]{[]int{1}}, 0, require.Error},
		{"no args", args[int]{}, 0, require.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LCM(tt.args.n...)
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAbs(t *testing.T) {
	type args[T constraints.Signed | constraints.Float] struct {
		x T
	}
	type testCase[T constraints.Signed | constraints.Float] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{"positive", args[int]{1}, 1},
		{"negative", args[int]{-1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Abs(tt.args.x))
		})
	}
}

func TestPow(t *testing.T) {
	type args[T constraints.Integer | constraints.Float] struct {
		x T
		y T
	}
	type testCase[T constraints.Integer | constraints.Float] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{"2^2", args[int]{2, 2}, 4},
		{"2^10", args[int]{2, 10}, 1024},
		{"4^4", args[int]{4, 4}, 256},
		{"4^10", args[int]{4, 10}, 1048576},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Pow(tt.args.x, tt.args.y))
		})
	}
}
