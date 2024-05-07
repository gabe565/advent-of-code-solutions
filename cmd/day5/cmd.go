package day5

import (
	"io"
	"slices"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
)

type Result struct {
	Individual int `toml:"individual"`
	Ranges     int `toml:"ranges"`
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
	b, err := io.ReadAll(cmd.InOrStdin())
	if err != nil {
		return err
	}

	var almanac Almanac
	if err := almanac.UnmarshalText(b); err != nil {
		return err
	}

	result := Result{
		Individual: slices.Min(almanac.Locations()),
		Ranges:     almanac.MinLocationRange(),
	}

	return toml.NewEncoder(cmd.OutOrStdout()).Encode(result)
}
