package day8

import (
	"math/bits"
	"strings"
)

func NewDigit(s string) Digit {
	var d Digit
	for _, b := range s {
		d |= 1 << (b - 'a')
	}
	return d
}

type Digit byte

func (d Digit) String() string {
	var str strings.Builder
	str.Grow(7)
	for i := range 7 {
		if d&(1<<i) != 0 {
			str.WriteByte(byte('a' + i))
		}
	}
	return str.String()
}

func (d Digit) Segments() int {
	return bits.OnesCount8(uint8(d))
}

func (d Digit) Decode() int {
	switch d.Segments() {
	case 2:
		return 1
	case 4:
		return 4
	case 3:
		return 7
	case 7:
		return 8
	default:
		return -1
	}
}
