package day5

import (
	"bufio"
	"bytes"
	"io"
	"strconv"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
	"golang.org/x/exp/slices"
)

func New() day.Day[*Printer, int] {
	return day.Day[*Printer, int]{
		Year: 2024,
		Day:  5,
		Parse: func(r io.Reader) (*Printer, error) {
			var p Printer
			scanner := bufio.NewScanner(r)
			ruleSection := true
			for scanner.Scan() {
				if len(scanner.Bytes()) == 0 {
					ruleSection = false
					continue
				}

				sep := []byte{'|'}
				if !ruleSection {
					sep = []byte{','}
				}

				vals := make([]int, 0, bytes.Count(scanner.Bytes(), sep))
				for _, val := range bytes.Split(scanner.Bytes(), sep) {
					parsed, err := strconv.Atoi(string(val))
					if err != nil {
						return nil, err
					}
					vals = append(vals, parsed)
				}

				if ruleSection {
					p.Rules = append(p.Rules, vals)
				} else {
					p.Updates = append(p.Updates, vals)
				}
			}
			return &p, scanner.Err()
		},
		Part1: func(printer *Printer) (int, error) {
			var result int
			for _, update := range printer.Updates {
				if printer.Valid(update) {
					result += update[len(update)/2]
				}
			}
			return result, nil
		},
		Part2: func(printer *Printer) (int, error) {
			var result int
			for _, update := range printer.Updates {
				if !printer.Valid(update) {
					update = printer.Fix(update)
					result += update[len(update)/2]
				}
			}
			return result, nil
		},
	}
}

type Printer struct {
	Rules   [][]int
	Updates [][]int
}

func (p *Printer) Valid(update []int) bool {
	for _, rule := range p.Rules {
		firstIndex := slices.Index(update, rule[0])
		if firstIndex == -1 {
			continue
		}
		secondIndex := slices.Index(update, rule[1])
		if secondIndex == -1 {
			continue
		}
		if firstIndex > secondIndex {
			return false
		}
	}
	return true
}

func (p *Printer) Fix(update []int) []int {
	var changed bool
	for _, rule := range p.Rules {
		firstIndex := slices.Index(update, rule[0])
		if firstIndex == -1 {
			continue
		}
		secondIndex := slices.Index(update, rule[1])
		if secondIndex == -1 {
			continue
		}
		if firstIndex > secondIndex {
			changed = true
			update[firstIndex], update[secondIndex] = update[secondIndex], update[firstIndex]
		}
	}
	if changed {
		update = p.Fix(update)
	}
	return update
}
