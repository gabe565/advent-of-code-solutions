package day9

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
		Use:   "day9",
		Short: "Day 9",
		RunE:  run,
	}
	return cmd
}

func run(cmd *cobra.Command, _ []string) error {
	var report Report
	if err := report.Decode(cmd.InOrStdin()); err != nil {
		return err
	}

	result := Result{
		Part1: report.Predict(PredictFuture),
		Part2: report.Predict(PredictPast),
	}

	return toml.NewEncoder(cmd.OutOrStdout()).Encode(result)
}
