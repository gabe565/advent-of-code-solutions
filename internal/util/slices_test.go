package util

import (
	"cmp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Parallel()
	type args[E cmp.Ordered] struct {
		s []E
	}
	type testCase[E cmp.Ordered] struct {
		name string
		args args[E]
		want E
	}
	tests := []testCase[int]{
		{"empty", args[int]{[]int{}}, 0},
		{"simple", args[int]{[]int{1, 2, 3}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, Sum(tt.args.s))
		})
	}
}
