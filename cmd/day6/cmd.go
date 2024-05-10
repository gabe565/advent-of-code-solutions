package day6

import (
	"io"

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
	b, err := io.ReadAll(cmd.InOrStdin())
	if err != nil {
		return err
	}

	var race Race
	if err := race.UnmarshalText(b); err != nil {
		return err
	}

	var result Result
	if result.Part1, err = race.Part1(); err != nil {
		return err
	}
	if result.Part2, err = race.Part2(); err != nil {
		return err
	}

	return toml.NewEncoder(cmd.OutOrStdout()).Encode(result)
}
