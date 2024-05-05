package cmd

import "github.com/spf13/cobra"

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use: "advent-of-code-2023",
	}
	return cmd
}
