package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLine_Direction(t *testing.T) {
	tests := []struct {
		name string
		line Line
		want Direction
	}{
		{"point", Ln(0, 0, 0, 0), Point},
		{"horizontal", Ln(0, 0, 1, 0), Horizontal},
		{"vertical", Ln(0, 0, 0, 1), Vertical},
		{"diagonal", Ln(0, 0, 1, 1), Diagonal},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.line.Direction())
		})
	}
}
