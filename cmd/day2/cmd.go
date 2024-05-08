package day2

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
		Use:   "day2",
		Short: "Day 2",
		RunE:  run,
	}
	return cmd
}

func run(cmd *cobra.Command, _ []string) error {
	var games []Game
	scan := bufio.NewScanner(cmd.InOrStdin())
	for i := 0; scan.Scan(); i++ {
		var game Game
		if err := game.UnmarshalText(scan.Bytes()); err != nil {
			return fmt.Errorf("failed to unmarshal game %d: %w", i, err)
		}
		games = append(games, game)
	}
	if scan.Err() != nil {
		return scan.Err()
	}

	var result Result
	for _, game := range games {
		result.Part2 += game.Power()
		if game.Valid() {
			result.Part1 += game.ID
		}
	}

	return toml.NewEncoder(cmd.OutOrStdout()).Encode(result)
}
