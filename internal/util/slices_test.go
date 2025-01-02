package util

import (
	"cmp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	type args[T cmp.Ordered] struct {
		s []T
	}
	type testCase[T cmp.Ordered] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{"empty", args[int]{[]int{}}, 0},
		{"simple", args[int]{[]int{1, 2, 3}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Sum(tt.args.s))
		})
	}
}

func TestStringToIntSlice(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr require.ErrorAssertionFunc
	}{
		{"empty", args{"", ""}, []int{}, require.NoError},
		{"single digits", args{"1 2 3", " "}, []int{1, 2, 3}, require.NoError},
		{"extra spaces", args{" 1  2 3  ", " "}, []int{1, 2, 3}, require.NoError},
		{"double digits", args{"10 20 30", " "}, []int{10, 20, 30}, require.NoError},
		{"mixed digits", args{"10  1  20", " "}, []int{10, 1, 20}, require.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToIntSlice(tt.args.s, tt.args.sep)
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
