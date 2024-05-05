package cmd

import (
	"github.com/gabe565/advent-of-code-2023/cmd/day1"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use: "advent-of-code-2023",
	}
	cmd.AddCommand(
		day1.New(),
	)
	return cmd
}
