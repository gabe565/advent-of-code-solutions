package year2021

import (
	"github.com/gabe565/advent-of-code-solutions/cmd/year2021/day1"
	"github.com/gabe565/advent-of-code-solutions/cmd/year2021/day2"
	"github.com/gabe565/advent-of-code-solutions/cmd/year2021/day3"
	"github.com/gabe565/advent-of-code-solutions/cmd/year2021/day4"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "2021",
		Short: "Solutions for 2021",
	}
	cmd.AddCommand(
		day1.New().Cmd(),
		day2.New().Cmd(),
		day3.New().Cmd(),
		day4.New().Cmd(),
	)
	return cmd
}
