package day4

import (
	"github.com/gabe565/advent-of-code-2023/internal/util"
)

type Cards struct {
	Cards []Card
}

func (c *Cards) Winning() int {
	counts := make([]int, len(c.Cards))
	for i, card := range c.Cards {
		counts[i]++
		for j := range card.Matches() {
			k := 1 + i + j
			counts[k] += counts[i]
		}
	}
	return util.Sum(counts)
}
