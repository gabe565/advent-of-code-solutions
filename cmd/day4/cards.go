package day4

import (
	"bufio"
	"fmt"
	"io"

	"github.com/gabe565/advent-of-code-2023/internal/util"
)

type Cards []Card

func (c *Cards) Decode(r io.Reader) error {
	scan := bufio.NewScanner(r)
	for i := 0; scan.Scan(); i++ {
		var card Card
		if err := card.UnmarshalText(scan.Bytes()); err != nil {
			return fmt.Errorf("failed to unmarshal card %d: %w", i, err)
		}
		*c = append(*c, card)
	}
	return scan.Err()
}

func (c Cards) Winning() int {
	counts := make([]int, len(c))
	for i, card := range c {
		counts[i]++
		for j := range card.Matches() {
			k := 1 + i + j
			counts[k] += counts[i]
		}
	}
	return util.Sum(counts)
}
