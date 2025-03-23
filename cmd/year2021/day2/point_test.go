package day2

import (
	"io"
	"strings"
	"testing"
)

func TestPoint_Write(t *testing.T) {
	type fields struct {
		X    int
		Y    int
		Aim  int
		Mode PointMode
	}
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		{"", fields{}, args{[]byte("forward 5")}, 9, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			point := &Point{
				X:    tt.fields.X,
				Y:    tt.fields.Y,
				Aim:  tt.fields.Aim,
				Mode: tt.fields.Mode,
			}
			gotN, err := point.Write(tt.args.p)
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

func TestPoint_ReadFrom(t *testing.T) {
	type fields struct {
		X    int
		Y    int
		Aim  int
		Mode PointMode
	}
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int64
		wantErr bool
	}{
		{"", fields{}, args{strings.NewReader("forward 5")}, 9, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			point := &Point{
				X:    tt.fields.X,
				Y:    tt.fields.Y,
				Aim:  tt.fields.Aim,
				Mode: tt.fields.Mode,
			}
			gotN, err := point.ReadFrom(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("ReadFrom() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestPoint_Move(t *testing.T) {
	type fields struct {
		X    int
		Y    int
		Aim  int
		Mode PointMode
	}
	type args struct {
		direction string
		amount    int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Point
		wantErr bool
	}{
		{"Linear", fields{Aim: 1}, args{DirForward, 1}, Point{X: 1, Aim: 1}, false},
		{"Aim", fields{Mode: MoveAim, Aim: 1}, args{DirForward, 1}, Point{X: 1, Y: 1, Aim: 1, Mode: MoveAim}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			point := &Point{
				X:    tt.fields.X,
				Y:    tt.fields.Y,
				Aim:  tt.fields.Aim,
				Mode: tt.fields.Mode,
			}
			if err := point.Move(tt.args.direction, tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("Move() error = %v, wantErr %v", err, tt.wantErr)
			}
			if *point != tt.want {
				t.Errorf("Move(), got = %v, want = %v", point, tt.want)
				return
			}
		})
	}
}

func TestPoint_MoveLinear(t *testing.T) {
	type fields struct {
		X   int
		Y   int
		Aim int
	}
	type args struct {
		direction string
		amount    int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Point
		wantErr bool
	}{
		{DirForward, fields{}, args{DirForward, 1}, Point{X: 1}, false},
		{DirDown, fields{}, args{DirDown, 1}, Point{Y: 1}, false},
		{DirUp, fields{}, args{DirUp, 1}, Point{Y: -1}, false},
		{"invalid", fields{}, args{"", 1}, Point{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			point := &Point{
				X:   tt.fields.X,
				Y:   tt.fields.Y,
				Aim: tt.fields.Aim,
			}
			err := point.moveLinear(tt.args.direction, tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("MoveLinear() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if *point != tt.want {
				t.Errorf("MoveLinear(), got = %v, want = %v", point, tt.want)
				return
			}
		})
	}
}

func TestPoint_MoveAim(t *testing.T) {
	type fields struct {
		X   int
		Y   int
		Aim int
	}
	type args struct {
		direction string
		amount    int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Point
		wantErr bool
	}{
		{DirForward, fields{}, args{DirForward, 1}, Point{X: 1}, false},
		{DirForward + DirDown, fields{Aim: 1}, args{DirForward, 1}, Point{X: 1, Y: 1, Aim: 1}, false},
		{DirDown, fields{}, args{DirDown, 1}, Point{Aim: 1}, false},
		{DirUp, fields{}, args{DirUp, 1}, Point{Aim: -1}, false},
		{"invalid", fields{}, args{"", 1}, Point{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			point := &Point{
				X:   tt.fields.X,
				Y:   tt.fields.Y,
				Aim: tt.fields.Aim,
			}
			err := point.moveAim(tt.args.direction, tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("MoveAim() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if *point != tt.want {
				t.Errorf("MoveAim(), got = %v, want = %v", point, tt.want)
				return
			}
		})
	}
}

func TestPoint_Multiply(t *testing.T) {
	type fields struct {
		X   int
		Y   int
		Aim int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"0,1", fields{X: 0, Y: 1}, 0},
		{"1,0", fields{X: 1, Y: 0}, 0},
		{"1,1", fields{X: 1, Y: 1}, 1},
		{"2,2", fields{X: 2, Y: 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Point{
				X:   tt.fields.X,
				Y:   tt.fields.Y,
				Aim: tt.fields.Aim,
			}
			if got := p.Multiply(); got != tt.want {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
