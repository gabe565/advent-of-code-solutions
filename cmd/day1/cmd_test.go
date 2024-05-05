package day1

import (
	"bytes"
	_ "embed"
	"strconv"
	"testing"

	"github.com/pelletier/go-toml/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed example_p1.txt
var exampleP1 []byte

//go:embed input.txt
var input []byte

//go:embed example_p2.txt
var exampleP2 []byte

func Test_run(t *testing.T) {
	t.Parallel()
	type args struct {
		input   []byte
		spelled bool
	}
	tests := []struct {
		name    string
		args    args
		want    Result
		wantErr require.ErrorAssertionFunc
	}{
		{"example p1", args{exampleP1, false}, Result{142}, require.NoError},
		{"input p1", args{input, false}, Result{54990}, require.NoError},
		{"example p2", args{exampleP2, true}, Result{281}, require.NoError},
		{"input p2", args{input, true}, Result{54473}, require.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			cmd := New()
			cmd.SetIn(bytes.NewReader(tt.args.input))
			var buf bytes.Buffer
			cmd.SetOut(&buf)
			require.NoError(t, cmd.Flags().Set("spelled", strconv.FormatBool(tt.args.spelled)))

			tt.wantErr(t, run(cmd, nil))
			var got Result
			require.NoError(t, toml.Unmarshal(buf.Bytes(), &got))
			assert.Equal(t, tt.want, got)
		})
	}
}
