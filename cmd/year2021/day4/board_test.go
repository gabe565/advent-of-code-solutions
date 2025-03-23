package day4

import "testing"

func TestBoard_String(t *testing.T) {
	tests := []struct {
		name  string
		board Board
		want  string
	}{
		{"1 line", Board{[]Cell{{1, false}, {2, false}, {3, false}}}, " 1  2  3 "},
		{"2 lines", Board{[]Cell{{1, false}}, []Cell{{2, false}}}, " 1 \n 2 "},
		{"winner", Board{[]Cell{{1, true}}}, "\033[1m 1 \033[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.board.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_Sum(t *testing.T) {
	tests := []struct {
		name  string
		board Board
		want  int
	}{
		{"empty", Board{}, 0},
		{"not drawn", Board{[]Cell{{1, false}}}, 1},
		{"drawn", Board{[]Cell{{1, true}}}, 0},
		{"2 values", Board{[]Cell{{1, false}, {1, false}}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.board.Sum(); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_Wins(t *testing.T) {
	tests := []struct {
		name  string
		board Board
		want  bool
	}{
		{"horizontal", Board{[]Cell{{1, true}, {2, true}}, []Cell{{3, false}, {4, false}}}, true},
		{"vertical", Board{[]Cell{{1, true}, {2, false}}, []Cell{{3, true}, {4, false}}}, true},
		{"false", Board{[]Cell{{1, true}, {2, false}}, []Cell{{3, false}, {4, false}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.board.Wins(); got != tt.want {
				t.Errorf("Wins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_Write(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		board   Board
		args    args
		wantN   int
		wantErr bool
	}{
		{"1 row", Board{}, args{[]byte("1 2")}, 2, false},
		{"2 rows", Board{}, args{[]byte("1\n2")}, 2, false},
		{"invalid", Board{}, args{[]byte("a")}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := tt.board.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("Write() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
