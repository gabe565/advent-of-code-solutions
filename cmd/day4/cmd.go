package day4

import (
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
	if err := cards.Decode(cmd.InOrStdin()); err != nil {
		return err
	}

	result := Result{Part2: cards.Winning()}
	for _, card := range cards {
		result.Part1 += card.Points()
	}

	return toml.NewEncoder(cmd.OutOrStdout()).Encode(result)
}
