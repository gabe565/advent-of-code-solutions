package day3

import (
	"bytes"
	"io"
	"regexp"
	"strconv"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
)

type Result struct {
	Part1 int `toml:"part1"`
	Part2 int `toml:"part2"`
}

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "day3",
		Short: "Day 3",
		RunE:  run,
	}
	return cmd
}

type Number struct {
	Value int
	X, Y  int
	Len   int
}

func run(cmd *cobra.Command, _ []string) error {
	var result Result

	input, err := io.ReadAll(cmd.InOrStdin())
	if err != nil {
		return err
	}
	schematic := bytes.Split(input, []byte("\n"))

	var numbers []Number
	splitRe := regexp.MustCompile("[^0-9]")
	for x, line := range schematic {
		split := splitRe.Split(string(line), -1)
		var y int
		for _, val := range split {
			if len(val) != 0 {
				parsed, err := strconv.Atoi(val)
				if err != nil {
					continue
				}
				numbers = append(numbers, Number{
					Value: parsed,
					X:     x,
					Y:     y,
					Len:   len(val),
				})
			}
			y += 1 + len(val)
		}
	}

	for x, line := range schematic {
		for y, char := range line {
			if char != '.' && char < '0' || char > '9' {
				computeRatio := char == '*'
				localMatches := make([]int, 0, 2)
				adjacent := [][]byte{
					schematic[x-1][y-1 : y+2],
					schematic[x][y-1 : y+2],
					schematic[x+1][y-1 : y+2],
				}
				for ax, line := range adjacent {
					var skip int
					for ay, char := range line {
						if skip > 0 {
							skip--
							continue
						}
						if char >= '0' && char <= '9' {
							realX := x - 1 + ax
							realY := y - 1 + ay
							for _, n := range numbers {
								if realX == n.X && realY >= n.Y && realY <= n.Y+n.Len {
									result.Part1 += n.Value
									skip = n.Y - realY + n.Len - 1
									if computeRatio {
										localMatches = append(localMatches, n.Value)
									}
								}
							}
							if len(localMatches) == 2 {
								result.Part2 += localMatches[0] * localMatches[1]
							}
						}
					}
				}
			}
		}
	}

	return toml.NewEncoder(cmd.OutOrStdout()).Encode(result)
}
