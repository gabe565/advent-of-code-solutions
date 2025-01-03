package year2024

import (
	"github.com/gabe565/advent-of-code-solutions/cmd/year2024/day1"
	"github.com/gabe565/advent-of-code-solutions/cmd/year2024/day2"
	"github.com/gabe565/advent-of-code-solutions/cmd/year2024/day3"
	"github.com/gabe565/advent-of-code-solutions/cmd/year2024/day4"
	"github.com/gabe565/advent-of-code-solutions/cmd/year2024/day5"
	"github.com/gabe565/advent-of-code-solutions/cmd/year2024/day6"
	"github.com/gabe565/advent-of-code-solutions/cmd/year2024/day7"
	"github.com/gabe565/advent-of-code-solutions/cmd/year2024/day8"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "2024",
		Short: "Solutions for 2024",
	}
	cmd.AddCommand(
		day1.New().Cmd(),
		day2.New().Cmd(),
		day3.New().Cmd(),
		day4.New().Cmd(),
		day5.New().Cmd(),
		day6.New().Cmd(),
		day7.New().Cmd(),
		day8.New().Cmd(),
	)
	return cmd
}
