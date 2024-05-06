package day2

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
)

type Result struct {
	Sum   int `toml:"sum"`
	Power int `toml:"power"`
}

const (
	RedCubes   = 12
	GreenCubes = 13
	BlueCubes  = 14
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "day2",
		Short: "Day 2",
		RunE:  run,
	}
	return cmd
}

func run(cmd *cobra.Command, _ []string) error {
	var result Result
	scan := bufio.NewScanner(cmd.InOrStdin())
	for scan.Scan() {
		line := scan.Text()
		var r, g, b int

		gameSpec, roundsSpec, ok := strings.Cut(line, ":")
		if !ok {
			return fmt.Errorf("invalid game: %s", line)
		}
		roundsSpec = strings.TrimSpace(roundsSpec)
		for _, round := range strings.Split(roundsSpec, ";") {
			for _, colors := range strings.Split(round, ",") {
				var color string
				var number int
				if _, err := fmt.Sscanf(colors, "%d %s", &number, &color); err != nil {
					return err
				}

				switch color {
				case "red":
					if r < number {
						r = number
					}
				case "green":
					if g < number {
						g = number
					}
				case "blue":
					if b < number {
						b = number
					}
				default:
					return fmt.Errorf("invalid color: %s", color)
				}
			}
		}

		if r <= RedCubes && g <= GreenCubes && b <= BlueCubes {
			var id int
			if _, err := fmt.Sscanf(gameSpec, "Game %d", &id); err != nil {
				return err
			}
			result.Sum += id
		}
		result.Power += r * g * b
	}
	if scan.Err() != nil {
		return scan.Err()
	}

	return toml.NewEncoder(cmd.OutOrStdout()).Encode(result)
}
