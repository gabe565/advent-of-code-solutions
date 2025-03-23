package day3

import (
	"reflect"
	"testing"
)

func TestNewDiagnostic(t *testing.T) {
	tests := []struct {
		name string
		want *Diagnostic
	}{
		{"Empty", &Diagnostic{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDiagnostic(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDiagnostic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLifeSupportDiagnostic(t *testing.T) {
	tests := []struct {
		name string
		want *Diagnostic
	}{
		{"Empty", &Diagnostic{storeEntries: true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLifeSupportDiagnostic(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDiagnostic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiagnostic_Add(t *testing.T) {
	type fields struct {
		sums []int
		x, y int
	}
	type args struct {
		v []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Example", fields{make([]int, 5), 0, 0}, args{[]byte("00100")}, false},
		{"Empty", fields{make([]int, 5), 0, 0}, args{[]byte("")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Diagnostic{
				sums: tt.fields.sums,
				x:    tt.fields.x,
				y:    tt.fields.y,
			}
			if err := d.Add(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDiagnostic_Epsilon(t *testing.T) {
	type fields struct {
		sums []int
		x, y int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"Example", fields{[]int{7, 5, 8, 7, 5}, 5, 12}, 9},
		{"Empty", fields{make([]int, 5), 5, 12}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Diagnostic{
				sums: tt.fields.sums,
				x:    tt.fields.x,
				y:    tt.fields.y,
			}
			if got := d.Epsilon(); got != tt.want {
				t.Errorf("Epsilon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiagnostic_Gamma(t *testing.T) {
	type fields struct {
		sums []int
		x, y int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"Example", fields{[]int{7, 5, 8, 7, 5}, 5, 12}, 22},
		{"Empty", fields{make([]int, 5), 0, 12}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Diagnostic{
				sums: tt.fields.sums,
				x:    tt.fields.x,
				y:    tt.fields.y,
			}
			if got := d.Gamma(); got != tt.want {
				t.Errorf("Gamma() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiagnostic_PowerConsumption(t *testing.T) {
	type fields struct {
		sums []int
		x, y int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"Example", fields{[]int{7, 5, 8, 7, 5}, 5, 12}, 198},
		{"Empty", fields{make([]int, 5), 0, 12}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Diagnostic{
				sums: tt.fields.sums,
				x:    tt.fields.x,
				y:    tt.fields.y,
			}
			if got := d.PowerConsumption(); got != tt.want {
				t.Errorf("PowerConsumption() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiagnostic_O2(t *testing.T) {
	type fields struct {
		storeEntries bool
		entries      [][]byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{"", fields{true, [][]byte{[]byte("00100"), []byte("11110")}}, 30, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Diagnostic{
				storeEntries: tt.fields.storeEntries,
				entries:      tt.fields.entries,
			}
			got, err := d.O2()
			if (err != nil) != tt.wantErr {
				t.Errorf("O2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("O2() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiagnostic_CO2(t *testing.T) {
	type fields struct {
		storeEntries bool
		entries      [][]byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{"", fields{true, [][]byte{[]byte("00100"), []byte("11110")}}, 4, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Diagnostic{
				storeEntries: tt.fields.storeEntries,
				entries:      tt.fields.entries,
			}
			got, err := d.CO2()
			if (err != nil) != tt.wantErr {
				t.Errorf("CO2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CO2() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiagnostic_LifeSupport(t *testing.T) {
	type fields struct {
		storeEntries bool
		entries      [][]byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{"", fields{true, [][]byte{[]byte("00100"), []byte("11110")}}, 120, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Diagnostic{
				storeEntries: tt.fields.storeEntries,
				entries:      tt.fields.entries,
			}
			got, err := d.LifeSupport()
			if (err != nil) != tt.wantErr {
				t.Errorf("LifeSupport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LifeSupport() got = %v, want %v", got, tt.want)
			}
		})
	}
}
