package day7

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() day.Day[[]Calibration, int] {
	return day.Day[[]Calibration, int]{
		Year: 2024,
		Day:  7,
		Parse: func(r io.Reader) ([]Calibration, error) {
			scanner := bufio.NewScanner(r)
			var calibrations []Calibration
			for scanner.Scan() {
				strs := strings.Split(scanner.Text(), " ")

				var calibration Calibration
				var err error

				calibration.Result, err = strconv.Atoi(strings.TrimSuffix(strs[0], ":"))
				if err != nil {
					return nil, err
				}

				calibration.Numbers = make([]int, 0, len(strs)-1)
				for _, val := range strs[1:] {
					val, err := strconv.Atoi(val)
					if err != nil {
						return nil, err
					}
					calibration.Numbers = append(calibration.Numbers, val)
				}
				calibrations = append(calibrations, calibration)
			}
			return calibrations, scanner.Err()
		},
		Part1: func(input []Calibration) (int, error) {
			var result int
			for _, calibration := range input {
				if calibration.IsValid([]Operator{Add, Multiply}) {
					result += calibration.Result
				}
			}
			return result, nil
		},
		Part2: func(input []Calibration) (int, error) {
			var result int
			for _, calibration := range input {
				if calibration.IsValid([]Operator{Add, Multiply, Concat}) {
					result += calibration.Result
				}
			}
			return result, nil
		},
	}
}

type Calibration struct {
	Result  int
	Numbers []int
}

type Operator byte

const (
	Add      Operator = '+'
	Multiply Operator = '*'
	Concat   Operator = '|'
)

func (o Operator) Do(a, b int) int {
	switch o {
	case Add:
		return a + b
	case Multiply:
		return a * b
	case Concat:
		pad := 10
		for pad <= b {
			pad *= 10
		}
		return a*pad + b
	default:
		panic("invalid operator")
	}
}

func (c *Calibration) IsValid(operators []Operator) bool {
	return isValid(c.Result, c.Numbers[0], c.Numbers[1:], operators)
}

func isValid(target, result int, operands []int, operators []Operator) bool {
	for _, op := range operators {
		result := op.Do(result, operands[0])
		switch {
		case result > target:
		case len(operands) == 1:
			if target == result {
				return true
			}
		case isValid(target, result, operands[1:], operators):
			return true
		}
	}
	return false
}
