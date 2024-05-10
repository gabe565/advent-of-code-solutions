package day7

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
		Use:   "day7",
		Short: "Day 7",
		RunE:  run,
	}
	return cmd
}

func run(cmd *cobra.Command, _ []string) error {
	var game Game
	if err := game.Decode(cmd.InOrStdin()); err != nil {
		return err
	}

	result := Result{
		Part1: game.Winnings(false),
		Part2: game.Winnings(true),
	}

	return toml.NewEncoder(cmd.OutOrStdout()).Encode(result)
}
