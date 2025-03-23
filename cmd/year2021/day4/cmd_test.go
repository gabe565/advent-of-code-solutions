package day4

import (
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
	day := New()
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		cmd     *cobra.Command
		args    args
		want    int
		wantErr require.ErrorAssertionFunc
	}{
		{"example part 1", day.Part1Cmd(), args{"2021/4/example.txt"}, 4512, require.NoError},
		{"input part 1", day.Part1Cmd(), args{"2021/4/input.txt"}, 35711, require.NoError},
		{"example part 2", day.Part2Cmd(), args{"2021/4/example.txt"}, 1924, require.NoError},
		{"input part 2", day.Part2Cmd(), args{"2021/4/input.txt"}, 5586, require.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := inputs.FS.Open(tt.args.input)
			require.NoError(t, err)
			t.Cleanup(func() { _ = f.Close() })
			tt.cmd.SetIn(f)

			var buf strings.Builder
			tt.cmd.SetOut(&buf)

			tt.wantErr(t, tt.cmd.RunE(tt.cmd, nil))
			got, err := strconv.Atoi(strings.TrimSpace(buf.String()))
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
