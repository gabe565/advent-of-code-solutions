package day

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"

	"gabe565.com/utils/termx"
	"github.com/spf13/cobra"
)

type PartFunc[In, Out any] func(input In) (Out, error)

type Day[In, Out any] struct {
	Date  time.Time
	Parse func(r io.Reader) (In, error)
	Part1 PartFunc[In, Out]
	Part2 PartFunc[In, Out]
}

func (d Day[In, Out]) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     d.Date.Format("02"),
		Short:   "Solutions for " + d.Date.Format(time.DateOnly),
		Args:    cobra.MaximumNArgs(1),
		Aliases: []string{d.Date.Format("2")},
	}
	if d.Part1 != nil {
		cmd.AddCommand(d.Part1Cmd())
	}
	if d.Part2 != nil {
		cmd.AddCommand(d.Part2Cmd())
	}
	return cmd
}

func (d Day[In, Out]) Part1Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "1 [input]",
		Short: "Solution for part 1",
		RunE:  d.run(d.Part1),
	}
}

func (d Day[In, Out]) Part2Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "2 [input]",
		Short: "Solution for part 2",
		RunE:  d.run(d.Part2),
	}
}

func (d Day[In, Out]) run(partFunc PartFunc[In, Out]) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var r io.Reader
		switch {
		case len(args) == 1:
			f, err := os.Open(args[0])
			if err != nil {
				return err
			}
			defer func() {
				_ = f.Close()
			}()

			r = f
		case termx.IsTerminal(cmd.InOrStdin()):
			return cmd.Usage()
		default:
			r = io.NopCloser(cmd.InOrStdin())
		}

		start := time.Now()
		input, err := d.Parse(r)
		if err != nil {
			return err
		}
		slog.Info("Parsed input", "took", time.Since(start))
		if f, ok := r.(io.Closer); ok {
			_ = f.Close()
		}

		start = time.Now()
		result, err := partFunc(input)
		if err != nil {
			return err
		}
		slog.Info("Computed result", "took", time.Since(start))

		_, err = fmt.Fprintln(cmd.OutOrStdout(), result)
		return err
	}
}
