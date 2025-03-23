package day

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"strconv"
	"time"

	"gabe565.com/utils/termx"
	"github.com/spf13/cobra"
)

type PartFunc[In, Out any] func(input In) (Out, error)

type Day[In, Out any] struct {
	Year  int
	Day   int
	Parse func(r io.Reader) (In, error)
	Part1 PartFunc[In, Out]
	Part2 PartFunc[In, Out]
}

func (d Day[In, Out]) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   fmt.Sprintf("%02d", d.Day),
		Short: "Solutions for " + d.DateString(),
		Args:  cobra.MaximumNArgs(1),
	}
	if d.Day < 10 {
		cmd.Aliases = []string{strconv.Itoa(d.Day)}
	}
	cmd.AddGroup(&cobra.Group{
		ID:    "parts",
		Title: "Parts",
	})
	if d.Part1 != nil {
		cmd.AddCommand(d.Part1Cmd())
	}
	if d.Part2 != nil {
		cmd.AddCommand(d.Part2Cmd())
	}
	return cmd
}

func (d Day[In, Out]) DateString() string {
	return fmt.Sprintf("%04d-12-%02d", d.Year, d.Day)
}

func (d Day[In, Out]) Part1Cmd() *cobra.Command {
	return &cobra.Command{
		Use:     "1 [input]",
		Short:   "Solution for " + d.DateString() + " part 1",
		RunE:    d.run(d.Part1),
		GroupID: "parts",
	}
}

func (d Day[In, Out]) Part2Cmd() *cobra.Command {
	return &cobra.Command{
		Use:     "2 [input]",
		Short:   "Solution for " + d.DateString() + " part 2",
		RunE:    d.run(d.Part2),
		GroupID: "parts",
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
