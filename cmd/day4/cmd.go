package day4

import (
	"bufio"
	"fmt"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
)

type Result struct {
	Part1 int `toml:"part1"`
	Part2 int `toml:"part2"`
}

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "day4",
		Short: "Day 4",
		RunE:  run,
	}
	return cmd
}

func run(cmd *cobra.Command, _ []string) error {
	var cards Cards
	scan := bufio.NewScanner(cmd.InOrStdin())
	for i := 0; scan.Scan(); i++ {
		var card Card
		if err := card.UnmarshalText(scan.Bytes()); err != nil {
			return fmt.Errorf("failed to unmarshal card %d: %w", i, err)
		}
		cards.Cards = append(cards.Cards, card)
	}
	if scan.Err() != nil {
		return scan.Err()
	}

	result := Result{Part2: cards.Winning()}
	for _, card := range cards.Cards {
		result.Part1 += card.Points()
	}

	return toml.NewEncoder(cmd.OutOrStdout()).Encode(result)
}
