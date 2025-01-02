package day5

import (
	"bufio"
	"bytes"
	"cmp"
	"fmt"
	"io"
	"slices"

	"github.com/gabe565/advent-of-code-solutions/internal/util"
)

type Almanac struct {
	Seeds []int
	Maps  []Map
}

func (a *Almanac) Decode(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	var err error
	var block []byte
	for i := 0; scanner.Scan(); i++ {
		if i == 0 {
			line := bytes.TrimPrefix(scanner.Bytes(), []byte("seeds: "))
			if a.Seeds, err = util.StringToIntSlice(string(line), " "); err != nil {
				return err
			}
		} else {
			if scanner.Text() == "" && len(block) != 0 {
				var m Map
				if err := m.UnmarshalText(block); err != nil {
					return err
				}
				a.Maps = append(a.Maps, m)
				block = block[:0]
			} else {
				if len(block) != 0 {
					block = append(block, '\n')
				}
				block = append(block, scanner.Bytes()...)
			}
		}
	}
	return scanner.Err()
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
	seeds := a.SeedRanges()
	for _, m := range a.Maps {
		for _, rule := range m.Rules {
			for i, seed := range seeds {
				if seed.End <= rule.Start || rule.End <= seed.Start {
					// Seed and rule are disjoint
					continue
				}

				if seed.Start <= rule.Start {
					// Seed overlaps rule start. Seed is split at rule start.
					seeds = append(seeds, SeedRange{
						Start: seed.Start,
						End:   rule.Start - 1,
					})
					seed.Start = rule.Start
				}

				if seed.End > rule.End {
					// Seed overlaps rule end. Seed is split at rule end.
					seeds = append(seeds, SeedRange{
						Start: rule.End,
						End:   seed.End - 1,
					})
					seed.End = rule.End
				}

				// Seed within rule. Seed is translated.
				seed.Start += rule.Diff
				seed.End += rule.Diff
				seeds[i] = seed
			}
		}
	}

	return slices.MinFunc(seeds, func(a, b SeedRange) int {
		return cmp.Compare(a.Start, b.Start)
	}).Start
}

func (a *Almanac) SeedRanges() []SeedRange {
	ranges := make([]SeedRange, 0, len(a.Seeds)/2)
	for i := 0; i < len(a.Seeds); i += 2 {
		ranges = append(ranges, SeedRange{
			Start: a.Seeds[i],
			End:   a.Seeds[i] + a.Seeds[i+1],
		})
	}
	return ranges
}

type SeedRange struct {
	Start, End int
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
	Start, End, Diff int
}

func (r *Rule) UnmarshalText(text []byte) error {
	var dest, count int
	_, err := fmt.Sscanf(string(text), "%d %d %d", &dest, &r.Start, &count)
	r.End = r.Start + count
	r.Diff = dest - r.Start
	return err
}

func (r *Rule) Transform(i int) int {
	if r.Start <= i && i < r.End {
		return i + r.Diff
	}
	return i
}
