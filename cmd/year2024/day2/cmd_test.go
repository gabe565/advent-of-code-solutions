package day2

import (
	"bytes"
	_ "embed"
	"strconv"
	"strings"
	"testing"

	"github.com/gabe565/advent-of-code-solutions/inputs"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	example, err := inputs.FS.ReadFile("2024/2/example.txt")
	require.NoError(t, err)

	input, err := inputs.FS.ReadFile("2024/2/input.txt")
	require.NoError(t, err)

	day := New()
	type args struct {
		input []byte
	}
	tests := []struct {
		name    string
		cmd     *cobra.Command
		args    args
		want    int
		wantErr require.ErrorAssertionFunc
	}{
		{"example part 1", day.Part1Cmd(), args{example}, 2, require.NoError},
		{"example part 2", day.Part2Cmd(), args{example}, 4, require.NoError},
		{"input part 1", day.Part1Cmd(), args{input}, 680, require.NoError},
		{"input part 2", day.Part2Cmd(), args{input}, 710, require.NoError},
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
