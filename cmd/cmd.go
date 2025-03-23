package cmd

import (
	"log/slog"
	"time"

	"gabe565.com/utils/termx"
	"github.com/gabe565/advent-of-code-solutions/cmd/year2021"
	"github.com/gabe565/advent-of-code-solutions/cmd/year2023"
	"github.com/gabe565/advent-of-code-solutions/cmd/year2024"
	"github.com/lmittmann/tint"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "advent-of-code-solutions",
		Short: "Advent Of Code Solutions by gabe565",
		PersistentPreRun: func(cmd *cobra.Command, _ []string) {
			slog.SetDefault(slog.New(tint.NewHandler(cmd.ErrOrStderr(), &tint.Options{
				Level:      slog.LevelInfo,
				TimeFormat: time.TimeOnly,
				NoColor:    !termx.IsColor(cmd.ErrOrStderr()),
			})))
		},
	}
	cmd.AddCommand(
		year2021.New(),
		year2023.New(),
		year2024.New(),
	)
	return cmd
}
