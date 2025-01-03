package util

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutations(t *testing.T) {
	type args[T any] struct {
		x []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want [][]T
	}
	tests := []testCase[int]{
		{"empty", args[int]{}, nil},
		{"1 2", args[int]{[]int{1, 2}}, [][]int{{1, 2}, {2, 1}}},
		{"1 2 3", args[int]{[]int{1, 2, 3}}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slices.Collect(Permutations(tt.args.x))
			assert.Equal(t, tt.want, got)
		})
	}
}
