package day4

import (
	"bytes"
	"fmt"
	"slices"

	"github.com/gabe565/advent-of-code-solutions/internal/util"
)

type Card struct {
	ID      int
	Winning []int
	Values  []int
}

func (c *Card) UnmarshalText(text []byte) error {
	idSpec, numbersSpec, ok := bytes.Cut(text, []byte(":"))
	if !ok {
		return fmt.Errorf("invalid card: %s", text)
	}

	if _, err := fmt.Fscanf(bytes.NewReader(idSpec), "Card %d", &c.ID); err != nil {
		return err
	}

	winningSpec, valuesSpec, ok := bytes.Cut(numbersSpec, []byte("|"))
	if !ok {
		return fmt.Errorf("invalid numbers spec: %s", numbersSpec)
	}

	var err error

	if c.Winning, err = util.StringToIntSlice(string(winningSpec), " "); err != nil {
		return err
	}

	if c.Values, err = util.StringToIntSlice(string(valuesSpec), " "); err != nil {
		return err
	}

	return nil
}

func (c *Card) Matches() int {
	var count int
	for _, v := range c.Values {
		if slices.Contains(c.Winning, v) {
			count++
		}
	}
	return count
}

func (c *Card) Points() int {
	count := c.Matches()
	if count == 0 {
		return 0
	}
	result := 1
	for range count - 1 {
		result *= 2
	}
	return result
}
