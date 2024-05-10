package day5

import (
	"slices"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
)

type Result struct {
	Part1 int `toml:"part1"`
	Part2 int `toml:"part2"`
}

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "day5",
		Short: "Day 5",
		RunE:  run,
	}
	return cmd
}

func run(cmd *cobra.Command, _ []string) error {
	var almanac Almanac
	if err := almanac.Decode(cmd.InOrStdin()); err != nil {
		return err
	}

	result := Result{
		Part1: slices.Min(almanac.Locations()),
		Part2: almanac.MinLocationRange(),
	}

	return toml.NewEncoder(cmd.OutOrStdout()).Encode(result)
}
