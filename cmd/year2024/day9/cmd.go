package day9

import (
	"bytes"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() day.Day[Disk, int] {
	return day.Day[Disk, int]{
		Date: time.Date(2024, 12, 9, 0, 0, 0, 0, time.Local),
		Parse: func(r io.Reader) (Disk, error) {
			input, err := io.ReadAll(r)
			if err != nil {
				return nil, err
			}
			input = bytes.TrimSpace(input)
			result := make(Disk, 0, len(input)*5)
			var fileIndex int
			for i, b := range input {
				isFile := i%2 == 0

				var repeatChar *int
				if isFile {
					temp := fileIndex
					repeatChar = &temp
					fileIndex++
				}

				repeat, err := strconv.Atoi(string(b))
				if err != nil {
					return nil, err
				}

				for range repeat {
					result = append(result, repeatChar)
				}
			}
			return result, nil
		},
		Part1: func(input Disk) (int, error) {
			return input.Fragment().Checksum()
		},
		Part2: func(input Disk) (int, error) {
			return input.Compact().Checksum()
		},
	}
}

type Disk []*int

func (d Disk) Fragment() Disk {
	j := len(d) - 1
	for i := range d {
		switch {
		case d[i] != nil:
		case i >= j:
			return d
		default:
			for {
				if d[j] != nil {
					d[i], d[j] = d[j], d[i]
					j--
					break
				}
				j--
			}
		}
	}
	return d
}

func (d Disk) Compact() Disk {
	var (
		blockID    int
		blockStart = len(d) - 1
		blockEnd   int
	)
	// Find first block
	for i := len(d) - 1; i >= 0; i-- {
		if d[i] != nil {
			blockID = *d[i]
			blockStart, blockEnd = i, i
			break
		}
	}

	for blockID > 0 {
		// Find block start and end indexes
		var found bool
		for i := blockStart; i >= 0; i-- {
			if found {
				if d[i] == nil || *d[i] != blockID {
					blockStart = i + 1
					break
				}
			} else if d[i] != nil && *d[i] == blockID {
				found = true
				blockEnd = i
			}

			if i == 0 {
				panic("block not found: " + strconv.Itoa(blockID))
			}
		}

		blockSize := blockEnd - blockStart + 1

		// Find empty space large enough
		var emptySpace int
		for i := range blockStart {
			if d[i] == nil {
				emptySpace++
				if emptySpace == blockSize {
					// Swap empty space and block
					for j := range blockSize {
						d[i-j], d[blockEnd-j] = d[blockEnd-j], d[i-j]
					}
					break
				}
			} else {
				emptySpace = 0
			}
		}
		blockID--
	}
	return d
}

func (d Disk) Checksum() (int, error) {
	var checksum int
	for i, b := range d {
		if b != nil {
			checksum += i * *b
		}
	}
	return checksum, nil
}

func (d Disk) String() string {
	var s strings.Builder
	s.Grow(len(d))
	for _, b := range d {
		switch {
		case b == nil:
			s.WriteByte('.')
		case *b < 10:
			s.WriteByte(byte(*b) + '0')
		case *b < 16:
			s.WriteByte(byte(*b-10) + 'a')
		default:
			s.WriteByte('?')
		}
	}
	return s.String()
}
