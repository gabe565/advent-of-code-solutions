package day2

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"strconv"
)

type Point struct {
	X, Y, Aim int
	Mode      PointMode
}

const (
	DirForward = "forward"
	DirDown    = "down"
	DirUp      = "up"
)

var ErrInvalidDirection = errors.New("invalid direction")

type FieldType uint8

func (point *Point) Write(p []byte) (int, error) {
	var direction string
	var amount int
	var n int
	for i, line := range bytes.Fields(p) {
		if len(line) == 0 {
			continue
		}
		if i%2 == 0 {
			direction = string(line)
		} else {
			var err error
			if amount, err = strconv.Atoi(string(line)); err != nil {
				return n, err
			}
			if err = point.Move(direction, amount); err != nil {
				return n, err
			}
		}
		n += len(line) + 1
	}
	return len(p), nil
}

func (point *Point) ReadFrom(r io.Reader) (int64, error) {
	var n int64
	s := bufio.NewScanner(r)
	for s.Scan() {
		written, err := point.Write(s.Bytes())
		if err != nil {
			return n, err
		}
		n += int64(written)
	}
	return n, s.Err()
}

var ErrInvalidMode = errors.New("invalid movement mode")

func (point *Point) Move(direction string, amount int) error {
	switch point.Mode {
	case MoveLinear:
		return point.moveLinear(direction, amount)
	case MoveAim:
		return point.moveAim(direction, amount)
	}
	return fmt.Errorf("%d: %w", point.Mode, ErrInvalidMode)
}

func (point *Point) moveLinear(direction string, amount int) error {
	switch direction {
	case DirForward:
		point.X += amount
	case DirDown:
		point.Y += amount
	case DirUp:
		point.Y -= amount
	default:
		return fmt.Errorf("%s: %w", direction, ErrInvalidDirection)
	}
	slog.Debug("Move",
		"direction", direction,
		"amount", amount,
		"point", point,
	)
	return nil
}

func (point *Point) moveAim(direction string, amount int) error {
	switch direction {
	case DirForward:
		point.Y += point.Aim * amount
		point.X += amount
	case DirDown:
		point.Aim += amount
	case DirUp:
		point.Aim -= amount
	default:
		return fmt.Errorf("%s: %w", direction, ErrInvalidDirection)
	}
	slog.Debug("Move",
		"direction", direction,
		"amount", amount,
		"point", point,
	)
	return nil
}

func (point Point) Multiply() int {
	return point.X * point.Y
}
