package day5

import (
	"image"
	"iter"

	"github.com/gabe565/advent-of-code-solutions/internal/util"
)

func Ln(x1, y1, x2, y2 int) Line {
	return Line{
		P1: image.Pt(x1, y1),
		P2: image.Pt(x2, y2),
	}
}

type Line struct {
	P1, P2 image.Point
}

func (l Line) String() string {
	return l.P1.String() + "-" + l.P2.String()
}

func (l Line) Direction() Direction {
	switch {
	case l.P1 == l.P2:
		return Point
	case l.P1.X == l.P2.X:
		return Vertical
	case l.P1.Y == l.P2.Y:
		return Horizontal
	default:
		return Diagonal
	}
}

func (l Line) Points() iter.Seq[image.Point] {
	return func(yield func(image.Point) bool) {
		if !yield(l.P1) {
			return
		}

		d := l.step()

		for p := l.P1; p != l.P2; {
			p = p.Add(d)
			if !yield(p) {
				return
			}
		}
	}
}

func (l Line) step() image.Point {
	return image.Pt(
		util.Sign(l.P2.X-l.P1.X),
		util.Sign(l.P2.Y-l.P1.Y),
	)
}
