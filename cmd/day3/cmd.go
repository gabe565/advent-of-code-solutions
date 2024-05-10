package day3

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
		Use:   "day3",
		Short: "Day 3",
		RunE:  run,
	}
	return cmd
}

func run(cmd *cobra.Command, _ []string) error {
	var schematic Schematic
	if err := schematic.Decode(cmd.InOrStdin()); err != nil {
		return err
	}

	var result Result
	result.Part1, result.Part2 = schematic.Ratios()

	return toml.NewEncoder(cmd.OutOrStdout()).Encode(result)
}
