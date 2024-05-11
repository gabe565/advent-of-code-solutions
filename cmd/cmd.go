package cmd

import (
	"github.com/gabe565/advent-of-code-2023/cmd/day1"
	"github.com/gabe565/advent-of-code-2023/cmd/day2"
	"github.com/gabe565/advent-of-code-2023/cmd/day3"
	"github.com/gabe565/advent-of-code-2023/cmd/day4"
	"github.com/gabe565/advent-of-code-2023/cmd/day5"
	"github.com/gabe565/advent-of-code-2023/cmd/day6"
	"github.com/gabe565/advent-of-code-2023/cmd/day7"
	"github.com/gabe565/advent-of-code-2023/cmd/day8"
	"github.com/gabe565/advent-of-code-2023/cmd/day9"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use: "advent-of-code-2023",
	}
	cmd.AddCommand(
		day1.New(),
		day2.New(),
		day3.New(),
		day4.New(),
		day5.New(),
		day6.New(),
		day7.New(),
		day8.New(),
		day9.New(),
	)
	return cmd
}
