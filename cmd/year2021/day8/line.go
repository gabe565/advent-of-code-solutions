package day8

import (
	"strings"
)

func NewLine(in string) Line {
	var line Line
	var keyIdx, codeIdx int

	var codeSection bool
	for word := range strings.SplitSeq(in, " ") {
		if word == "|" {
			codeSection = true
			continue
		}

		if codeSection {
			line.Code[codeIdx] = NewDigit(word)
			codeIdx++
		} else {
			line.Key[keyIdx] = NewDigit(word)
			keyIdx++
		}
	}

	line.DeriveKey()

	return line
}

type Line struct {
	Key  [10]Digit
	Code [4]Digit
}

func (l Line) String() string {
	var s strings.Builder
	for _, d := range l.Key {
		s.WriteString(d.String())
		s.WriteByte(' ')
	}
	s.WriteString("| ")
	for _, d := range l.Code {
		s.WriteString(d.String())
		s.WriteByte(' ')
	}
	return s.String()
}

func (l *Line) DeriveKey() {
	var k [10]Digit

	// 1, 4, 7, 8
	for _, d := range l.Key {
		if n := d.Decode(); n != -1 {
			k[n] = d
		}
	}

	// 0, 6, 9
	for _, d := range l.Key {
		if d.Segments() == 6 {
			switch {
			case k[4]&^d == 0:
				k[9] = d
			case k[1]&^d == 0:
				k[0] = d
			default:
				k[6] = d
			}
		}
	}

	// 2, 3, 5
	for _, d := range l.Key {
		if d.Segments() == 5 {
			switch {
			case k[1]&^d == 0:
				k[3] = d
			case d&^k[6] == 0:
				k[5] = d
			default:
				k[2] = d
			}
		}
	}

	l.Key = k
}
