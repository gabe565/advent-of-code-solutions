package day4

import "testing"

func TestCell_String(t *testing.T) {
	type fields struct {
		Value int
		Drawn bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"not drawn", fields{1, false}, " 1 "},
		{"drawn", fields{1, true}, "\033[1m 1 \033[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Cell{
				Value: tt.fields.Value,
				Drawn: tt.fields.Drawn,
			}
			if got := c.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
