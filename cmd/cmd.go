package cmd

import (
	"github.com/gabe565/advent-of-code-solutions/cmd/year2023"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "advent-of-code-solutions",
		Short: "Advent Of Code Solutions by gabe565",
	}
	cmd.AddCommand(
		year2023.New(),
	)
	return cmd
}
