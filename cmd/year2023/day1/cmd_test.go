package day1

import (
	"bytes"
	_ "embed"
	"strconv"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed example_p1.txt
var exampleP1 []byte

//go:embed example_p2.txt
var exampleP2 []byte

//go:embed input.txt
var input []byte

func TestSolution(t *testing.T) {
	day := New()
	type args struct {
		input   []byte
		spelled bool
	}
	tests := []struct {
		name    string
		cmd     *cobra.Command
		args    args
		want    int
		wantErr require.ErrorAssertionFunc
	}{
		{"example part 1", day.Part1Cmd(), args{exampleP1, false}, 142, require.NoError},
		{"example part 2", day.Part2Cmd(), args{exampleP2, true}, 281, require.NoError},
		{"input part 1", day.Part1Cmd(), args{input, false}, 54990, require.NoError},
		{"input part 2", day.Part2Cmd(), args{input, false}, 54473, require.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cmd.SetIn(bytes.NewReader(tt.args.input))
			var buf bytes.Buffer
			tt.cmd.SetOut(&buf)

			tt.wantErr(t, tt.cmd.RunE(tt.cmd, nil))
			got, err := strconv.Atoi(strings.TrimSpace(buf.String()))
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
