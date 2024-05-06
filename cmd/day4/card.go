package day4

import (
	"bytes"
	"fmt"
	"slices"
	"strconv"
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

	winningSpec = bytes.TrimSpace(winningSpec)
	for _, v := range bytes.Split(winningSpec, []byte(" ")) {
		if len(v) != 0 {
			v, err := strconv.Atoi(string(v))
			if err != nil {
				return err
			}
			c.Winning = append(c.Winning, v)
		}
	}

	valuesSpec = bytes.TrimSpace(valuesSpec)
	for _, v := range bytes.Split(valuesSpec, []byte(" ")) {
		if len(v) != 0 {
			v, err := strconv.Atoi(string(v))
			if err != nil {
				return err
			}
			c.Values = append(c.Values, v)
		}
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
