package day4

import "testing"

func Test_order_String(t *testing.T) {
	tests := []struct {
		name  string
		order Order
		want  string
	}{
		{"single", Order{1}, "1"},
		{"multiple", Order{1, 2, 3}, "1,2,3"},
		{"empty", Order{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.order.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_order_Write(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		order   Order
		args    args
		wantN   int
		wantErr bool
	}{
		{"simple", Order{}, args{[]byte("1,2,3")}, 3, false},
		{"empty", Order{}, args{[]byte("")}, 0, false},
		{"invalid", Order{}, args{[]byte("a")}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := tt.order.Write(tt.args.p)
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
