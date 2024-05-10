package day6

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/pelletier/go-toml/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed example.txt
var example []byte

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
		{"example", args{example}, Result{288, 71503}, require.NoError},
		{"input", args{input}, Result{449550, 28360140}, require.NoError},
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
