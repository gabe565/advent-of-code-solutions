package day8

import (
	"log/slog"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
)

type Result struct {
	Part1 int `toml:"part1"`
	Part2 int `toml:"part2"`
}

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "day8",
		Short: "Day 8",
		RunE:  run,
	}
	return cmd
}

func run(cmd *cobra.Command, _ []string) error {
	var network Network
	if err := network.Decode(cmd.InOrStdin()); err != nil {
		return err
	}

	part1, err := network.Steps("AAA", false)
	if err != nil {
		slog.Warn("Part 1 failed", "err", err.Error())
	}
	part2, err := network.GhostSteps()
	if err != nil {
		slog.Warn("Part 2 failed", "err", err.Error())
	}

	result := Result{
		Part1: part1,
		Part2: part2,
	}

	return toml.NewEncoder(cmd.OutOrStdout()).Encode(result)
}
