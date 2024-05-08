package day1

import (
	"bufio"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
)

type Result struct {
	Part1 int `toml:"part1"`
	Part2 int `toml:"part2"`
}

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "day1",
		Short: "Day 1",
		RunE:  run,
	}
	return cmd
}

func run(cmd *cobra.Command, _ []string) error {
	replacer := strings.NewReplacer(
		"one", "o1e",
		"two", "t2o",
		"three", "t3e",
		"four", "f4r",
		"five", "f5e",
		"six", "s6x",
		"seven", "s7n",
		"eight", "e8t",
		"nine", "n9e",
	)

	var result Result
	scan := bufio.NewScanner(cmd.InOrStdin())
	for scan.Scan() {
		line := scan.Text()
		spelledLine := line
		for {
			replaced := replacer.Replace(spelledLine)
			if spelledLine == replaced {
				break
			}
			spelledLine = replaced
		}

		result.Part1 += findFirstLast(line)
		result.Part2 += findFirstLast(spelledLine)
	}
	if scan.Err() != nil {
		return scan.Err()
	}

	return toml.NewEncoder(cmd.OutOrStdout()).Encode(result)
}

func findFirstLast(line string) int {
	var first *int
	var second int
	for _, b := range []byte(line) {
		if b >= '0' && b <= '9' {
			i := int(b - '0')
			if first == nil {
				first = &i
			}
			second = i
		}
	}

	if first == nil {
		return 0
	}
	return 10*(*first) + second
}
