package day5

import (
	"bytes"
	"errors"
	"fmt"
	"math"

	"github.com/gabe565/advent-of-code-2023/internal/util"
)

type Almanac struct {
	Seeds []int
	Maps  []Map
}

func (a *Almanac) UnmarshalText(text []byte) error {
	blocks := bytes.Split(text, []byte("\n\n"))
	if len(blocks) == 0 {
		return errors.New("invalid input")
	}

	if !bytes.HasPrefix(blocks[0], []byte("seeds: ")) {
		return fmt.Errorf("invalid seeds line: %q", string(blocks[0]))
	}
	blocks[0] = bytes.TrimPrefix(blocks[0], []byte("seeds: "))
	var err error
	a.Seeds, err = util.StringToIntSlice(string(blocks[0]), " ")
	if err != nil {
		return err
	}

	for _, block := range blocks[1:] {
		var m Map
		if err := m.UnmarshalText(block); err != nil {
			return err
		}
		a.Maps = append(a.Maps, m)
	}
	return nil
}

func (a *Almanac) Transform(i int) int {
	for _, m := range a.Maps {
		i = m.Transform(i)
	}
	return i
}

func (a *Almanac) Locations() []int {
	result := make([]int, 0, len(a.Seeds))
	for _, seed := range a.Seeds {
		result = append(result, a.Transform(seed))
	}
	return result
}

func (a *Almanac) MinLocationRange() int {
	result := math.MaxInt64
	for _, seedRange := range a.SeedRanges() {
		for i := range seedRange.Len {
			if location := a.Transform(i + seedRange.Start); location < result {
				result = location
			}
		}
	}
	return result
}

func (a *Almanac) SeedRanges() []SeedRange {
	ranges := make([]SeedRange, 0, len(a.Seeds)/2)
	for i := 0; i < len(a.Seeds); i += 2 {
		ranges = append(ranges, SeedRange{a.Seeds[i], a.Seeds[i+1]})
	}
	return ranges
}

type SeedRange struct {
	Start int
	Len   int
}

type Map struct {
	Name  string
	Rules []Rule
}

func (m *Map) UnmarshalText(text []byte) error {
	lines := bytes.Split(text, []byte("\n"))
	if _, err := fmt.Sscanf(string(lines[0]), "%s map:", &m.Name); err != nil {
		return err
	}

	for _, line := range lines[1:] {
		if len(line) != 0 {
			var r Rule
			if err := r.UnmarshalText(line); err != nil {
				return err
			}
			m.Rules = append(m.Rules, r)
		}
	}

	return nil
}

func (m *Map) Transform(a int) int {
	for _, rule := range m.Rules {
		if b := rule.Transform(a); b != a {
			return b
		}
	}
	return a
}

type Rule struct {
	Destination int
	Source      int
	Len         int
}

func (r *Rule) UnmarshalText(text []byte) error {
	_, err := fmt.Sscanf(string(text), "%d %d %d", &r.Destination, &r.Source, &r.Len)
	return err
}

func (r *Rule) Transform(i int) int {
	if r.Source <= i && i < r.Source+r.Len {
		return i + r.Destination - r.Source
	}
	return i
}
