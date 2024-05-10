package day8

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/pelletier/go-toml/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed example_2step.txt
var example2Step []byte

//go:embed example_6step.txt
var example6Step []byte

//go:embed example_part2.txt
var examplePart2 []byte

//go:embed input.txt
var input []byte

func Test_run(t *testing.T) {
	t.Parallel()
	type args struct {
		input []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Result
		wantErr require.ErrorAssertionFunc
	}{
		{"example 2 steps", args{example2Step}, Result{2, 0}, require.NoError},
		{"example 6 steps", args{example6Step}, Result{6, 0}, require.NoError},
		{"example part 2", args{examplePart2}, Result{0, 6}, require.NoError},
		{"input", args{input}, Result{15989, 13830919117339}, require.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			cmd := New()
			cmd.SetIn(bytes.NewReader(tt.args.input))
			var buf bytes.Buffer
			cmd.SetOut(&buf)

			tt.wantErr(t, run(cmd, nil))
			var got Result
			require.NoError(t, toml.Unmarshal(buf.Bytes(), &got))
			assert.Equal(t, tt.want, got)
		})
	}
}
