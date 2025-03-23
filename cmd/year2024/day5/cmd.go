package day5

import (
	"bufio"
	"bytes"
	"io"
	"slices"
	"strconv"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() day.Day[*Printer, int] {
	return day.Day[*Printer, int]{
		Year: 2024,
		Day:  5,
		Parse: func(r io.Reader) (*Printer, error) {
			var p Printer
			scanner := bufio.NewScanner(r)
			ruleSection := true
			sep := []byte{'|'}
			for scanner.Scan() {
				if len(scanner.Bytes()) == 0 {
					ruleSection = false
					sep[0] = ','
					continue
				}

				vals := make([]int, 0, bytes.Count(scanner.Bytes(), sep)+1)
				for val := range bytes.SplitSeq(scanner.Bytes(), sep) {
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
			for i, update := range printer.Updates {
				if printer.Valid(i) {
					result += update[len(update)/2]
				}
			}
			return result, nil
		},
		Part2: func(printer *Printer) (int, error) {
			var result int
			for i, update := range printer.Updates {
				if printer.Fix(i) {
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

func (p *Printer) Valid(i int) bool {
	update := p.Updates[i]
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

func (p *Printer) Fix(i int) bool {
	update := p.Updates[i]
	matching := make([][]int, 0, len(p.Rules))
	for _, rule := range p.Rules {
		if slices.Contains(update, rule[0]) && slices.Contains(update, rule[1]) {
			matching = append(matching, rule)
		}
	}

	var changed bool
	for {
		var changedThisRun bool
		for _, rule := range matching {
			firstIndex := slices.Index(update, rule[0])
			if firstIndex == -1 {
				panic("rule did not match")
			}
			secondIndex := slices.Index(update, rule[1])
			if secondIndex == -1 {
				panic("rule did not match")
			}
			if firstIndex > secondIndex {
				changed, changedThisRun = true, true
				update[firstIndex], update[secondIndex] = update[secondIndex], update[firstIndex]
			}
		}
		if !changedThisRun {
			break
		}
	}
	return changed
}
