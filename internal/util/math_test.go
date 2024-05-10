package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/constraints"
)

func TestGCD(t *testing.T) {
	t.Parallel()
	type args[E constraints.Integer] struct {
		a E
		b E
	}
	type testCase[E constraints.Integer] struct {
		name string
		args args[E]
		want E
	}
	tests := []testCase[int]{
		{"2 4", args[int]{2, 4}, 2},
		{"3 4", args[int]{3, 4}, 1},
		{"30 40", args[int]{30, 40}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, GCD(tt.args.a, tt.args.b))
		})
	}
}

func TestLCM(t *testing.T) {
	t.Parallel()
	type args[E constraints.Integer] struct {
		n []E
	}
	type testCase[E constraints.Integer] struct {
		name    string
		args    args[E]
		want    E
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
			t.Parallel()
			got, err := LCM(tt.args.n...)
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
