package day8

import (
	"bufio"
	"bytes"
	"io"
	"strings"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() *day.Day[[]*Line, int] {
	return &day.Day[[]*Line, int]{
		Year: 2021,
		Day:  8,
		Parse: func(r io.Reader) ([]*Line, error) {
			scanner := bufio.NewScanner(r)
			var buf strings.Builder
			var input []*Line

			for scanner.Scan() {
				if buf.Len() == 0 && bytes.HasSuffix(scanner.Bytes(), []byte("|")) {
					buf.Write(scanner.Bytes())
					buf.WriteByte(' ')
					continue
				}

				buf.Write(scanner.Bytes())

				input = append(input, new(NewLine(buf.String())))

				buf.Reset()
			}

			return input, scanner.Err()
		},
		Part1: func(input []*Line) (int, error) {
			var n int
			for _, line := range input {
				for _, d := range line.Code[len(line.Code)-4:] {
					if d.Decode() != -1 {
						n++
					}
				}
			}
			return n, nil
		},
		Part2: func(input []*Line) (int, error) {
			var sum int
			for _, line := range input {
				var code int
				for _, d := range line.Code {
					for n, k := range line.Key {
						if d == k {
							code = code*10 + n
							break
						}
					}
				}
				sum += code
			}
			return sum, nil
		},
	}
}
