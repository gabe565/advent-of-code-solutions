package day6

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
		Use:   "day6",
		Short: "Day 6",
		RunE:  run,
	}
	return cmd
}

func run(cmd *cobra.Command, _ []string) error {
	var race Race
	if err := race.Decode(cmd.InOrStdin()); err != nil {
		return err
	}

	var result Result
	var err error
	if result.Part1, err = race.Part1(); err != nil {
		return err
	}
	if result.Part2, err = race.Part2(); err != nil {
		return err
	}

	return toml.NewEncoder(cmd.OutOrStdout()).Encode(result)
}
