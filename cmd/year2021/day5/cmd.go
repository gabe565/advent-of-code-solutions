package day5

import (
	"bufio"
	"fmt"
	"image"
	"io"
	"slices"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() *day.Day[[]Line, int] {
	return &day.Day[[]Line, int]{
		Year: 2021,
		Day:  5,
		Parse: func(r io.Reader) ([]Line, error) {
			var input []Line
			scanner := bufio.NewScanner(r)
			for scanner.Scan() {
				var x1, y1, x2, y2 int
				if _, err := fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &x1, &y1, &x2, &y2); err != nil {
					return nil, err
				}

				input = append(input, Ln(x1, y1, x2, y2))
			}
			return input, scanner.Err()
		},
		Part1: func(input []Line) (int, error) {
			input = slices.DeleteFunc(input, func(r Line) bool {
				return r.Direction() == Diagonal
			})

			return countOverlaps(input), nil
		},
		Part2: func(input []Line) (int, error) {
			return countOverlaps(input), nil
		},
	}
}

func countOverlaps(input []Line) int {
	var bounds image.Rectangle
	for _, l := range input {
		bounds = bounds.Union(image.Rectangle{
			Min: l.P1,
			Max: l.P1.Add(image.Pt(1, 1)),
		})
		bounds = bounds.Union(image.Rectangle{
			Min: l.P2,
			Max: l.P2.Add(image.Pt(1, 1)),
		})
	}

	w := bounds.Dx()
	grid := make([]uint8, w*bounds.Dy())
	var n int
	for _, l := range input {
		for pt := range l.Points() {
			pt = pt.Sub(bounds.Min)
			i := pt.Y*w + pt.X
			if grid[i] < 2 {
				grid[i]++
				if grid[i] == 2 {
					n++
				}
			}
		}
	}
	return n
}
