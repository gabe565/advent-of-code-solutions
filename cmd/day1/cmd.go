package day1

import (
	"bufio"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
)

type Result struct {
	Sum int `toml:"sum"`
}

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "day1",
		Short: "Day 1",
		RunE:  run,
	}
	cmd.Flags().Bool("spelled", false, "Include spelled numbers")
	return cmd
}

func run(cmd *cobra.Command, _ []string) error {
	spelled, err := cmd.Flags().GetBool("spelled")
	if err != nil {
		return err
	}
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
		if spelled {
			for {
				newline := replacer.Replace(line)
				if line == newline {
					break
				}
				line = newline
			}
		}

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

		if first != nil {
			result.Sum += 10*(*first) + second
		}
	}
	if scan.Err() != nil {
		return scan.Err()
	}

	return toml.NewEncoder(cmd.OutOrStdout()).Encode(result)
}
