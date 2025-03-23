package day3

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"sync"
)

func NewDiagnostic() *Diagnostic {
	return &Diagnostic{}
}

func NewLifeSupportDiagnostic() *Diagnostic {
	d := NewDiagnostic()
	d.storeEntries = true
	return d
}

type Diagnostic struct {
	sums []int
	x, y int

	storeEntries bool
	entries      [][]byte

	once sync.Once
}

func (d *Diagnostic) Write(p []byte) (int, error) {
	var n int
	for _, line := range bytes.Fields(p) {
		if len(line) == 0 {
			continue
		}
		if err := d.Add(line); err != nil {
			return n, err
		}
		n += len(line) + 1
	}
	return len(p), nil
}

func (d *Diagnostic) ReadFrom(r io.Reader) (int64, error) {
	var n int64
	s := bufio.NewScanner(r)
	for s.Scan() {
		written, err := d.Write([]byte(s.Text()))
		n += int64(written)
		if err != nil {
			return n, err
		}
	}
	return n, s.Err()
}

var ErrBadWidth = errors.New("line width does not match other lines")

func (d *Diagnostic) Add(line []byte) error {
	if len(line) == 0 {
		return io.ErrUnexpectedEOF
	}

	d.once.Do(func() {
		d.x = len(line)
		d.sums = make([]int, len(line))
		if d.storeEntries {
			d.entries = make([][]byte, 0, 16)
		}
	})

	if len(line) != d.x {
		return fmt.Errorf("%s: %w", line, ErrBadWidth)
	}

	for i, b := range line {
		d.sums[i] += int(b - '0')
	}
	d.y++
	if d.storeEntries {
		d.entries = append(d.entries, line)
	}

	slog.Debug("Add value",
		"input", string(line),
		"sums", d.sums,
		"count", d.y,
	)

	return nil
}

func (d *Diagnostic) Gamma() int {
	var result int
	for i, sum := range d.sums {
		if sum > d.y/2 {
			result += 1 << (d.x - 1 - i)
		}
	}
	return result
}

func (d *Diagnostic) Epsilon() int {
	return d.Gamma() ^ (1<<d.x - 1)
}

func (d *Diagnostic) PowerConsumption() int {
	return d.Gamma() * d.Epsilon()
}

func commonBitAt(p [][]byte, n int) int {
	var sum float64
	for _, b := range p {
		sum += float64(b[n] - '0')
	}
	if sum >= float64(len(p))/2 {
		return 1
	}
	return 0
}

func uncommonBitAt(p [][]byte, n int) int {
	return commonBitAt(p, n) ^ 1
}

var ErrLifeSupportDiagnosticDisabled = errors.New("life support diagnostic is not enabled")

func (d *Diagnostic) lifeSupportDiag(bitSelector func([][]byte, int) int) (int, error) {
	if !d.storeEntries {
		return 0, ErrLifeSupportDiagnosticDisabled
	}

	if len(d.entries) == 0 {
		return 0, os.ErrInvalid
	}

	entries := &d.entries
	for i := range len((*entries)[0]) {
		bit := bitSelector(*entries, i)
		temp := make([][]byte, 0, len(*entries))
		for _, line := range *entries {
			firstBit := int(line[i] - '0')
			if firstBit == bit {
				temp = append(temp, line)
			}
		}
		entries = &temp
		if len(*entries) <= 1 {
			break
		}
	}

	if len(*entries) == 0 {
		return 0, os.ErrInvalid
	}

	var result int
	width := len((*entries)[0]) - 1
	for i, b := range (*entries)[0] {
		result += int(b-'0') << (width - i)
	}

	return result, nil
}

func (d *Diagnostic) O2() (int, error) {
	return d.lifeSupportDiag(commonBitAt)
}

func (d *Diagnostic) CO2() (int, error) {
	return d.lifeSupportDiag(uncommonBitAt)
}

func (d *Diagnostic) LifeSupport() (int, error) {
	if !d.storeEntries {
		return 0, ErrLifeSupportDiagnosticDisabled
	}

	o2, err := d.O2()
	if err != nil {
		return 0, err
	}

	co2, err := d.CO2()
	if err != nil {
		return 0, err
	}

	return o2 * co2, nil
}
