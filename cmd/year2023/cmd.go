package year2023

import (
	"github.com/gabe565/advent-of-code-solutions/cmd/year2023/day1"
	"github.com/gabe565/advent-of-code-solutions/cmd/year2023/day2"
	"github.com/gabe565/advent-of-code-solutions/cmd/year2023/day3"
	"github.com/gabe565/advent-of-code-solutions/cmd/year2023/day4"
	"github.com/gabe565/advent-of-code-solutions/cmd/year2023/day5"
	"github.com/gabe565/advent-of-code-solutions/cmd/year2023/day6"
	"github.com/gabe565/advent-of-code-solutions/cmd/year2023/day7"
	"github.com/gabe565/advent-of-code-solutions/cmd/year2023/day8"
	"github.com/gabe565/advent-of-code-solutions/cmd/year2023/day9"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "2023",
		Short: "Solutions for 2023",
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
		day9.New().Cmd(),
	)
	return cmd
}
