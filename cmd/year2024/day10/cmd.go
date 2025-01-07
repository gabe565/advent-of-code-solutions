package day10

import (
	"bytes"
	"cmp"
	"image"
	"io"
	"slices"
	"strings"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() day.Day[Map, int] {
	return day.Day[Map, int]{
		Year: 2024,
		Day:  10,
		Parse: func(r io.Reader) (Map, error) {
			b, err := io.ReadAll(r)
			if err != nil {
				return nil, err
			}
			b = bytes.TrimSpace(b)
			return bytes.Split(b, []byte("\n")), nil
		},
		Part1: func(input Map) (int, error) {
			var result int
			trailheads := input.FindTrailheads()
			for _, trailhead := range trailheads {
				paths := input.FindDestinations(trailhead)
				result += len(paths)
			}
			return result, nil
		},
		Part2: func(input Map) (int, error) {
			var result int
			trailheads := input.FindTrailheads()
			for _, trailhead := range trailheads {
				rating := input.FindRating(trailhead)
				result += rating
			}
			return result, nil
		},
	}
}

type Map [][]byte

func (m Map) String() string {
	var size int
	for _, line := range m {
		size += len(line) + 1
	}
	var s strings.Builder
	s.Grow(size)
	for i, line := range m {
		if i != 0 {
			s.WriteByte('\n')
		}
		s.Write(line)
	}
	return s.String()
}

func (m Map) FindTrailheads() []image.Point {
	var trailheads []image.Point
	for y, line := range m {
		for x, b := range line {
			if b == '0' {
				trailheads = append(trailheads, image.Pt(x, y))
			}
		}
	}
	return trailheads
}

func (m Map) FindPaths(p image.Point) []image.Point {
	val := m[p.Y][p.X]
	if val == '9' {
		return []image.Point{p}
	}

	var paths []image.Point
	if checkNorth(m, p.X, p.Y, val) {
		paths = append(paths, m.FindPaths(p.Add(image.Pt(0, -1)))...)
	}
	if checkEast(m, p.X, p.Y, val) {
		paths = append(paths, m.FindPaths(p.Add(image.Pt(1, 0)))...)
	}
	if checkSouth(m, p.X, p.Y, val) {
		paths = append(paths, m.FindPaths(p.Add(image.Pt(0, 1)))...)
	}
	if checkWest(m, p.X, p.Y, val) {
		paths = append(paths, m.FindPaths(p.Add(image.Pt(-1, 0)))...)
	}

	return paths
}

func (m Map) FindDestinations(p image.Point) []image.Point {
	paths := m.FindPaths(p)
	slices.SortStableFunc(paths, func(a, b image.Point) int {
		return cmp.Compare(10*a.X+a.Y, 10*b.X+b.Y)
	})
	return slices.Compact(paths)
}

func (m Map) FindRating(p image.Point) int {
	paths := m.FindPaths(p)
	return len(paths)
}

func checkNorth(input [][]byte, x, y int, val byte) bool {
	return y >= 1 && input[y-1][x] == val+1
}

func checkEast(input [][]byte, x, y int, val byte) bool {
	return x+1 < len(input[y]) && input[y][x+1] == val+1
}

func checkSouth(input [][]byte, x, y int, val byte) bool {
	return y+1 < len(input) && input[y+1][x] == val+1
}

func checkWest(input [][]byte, x, y int, val byte) bool {
	return x >= 1 && input[y][x-1] == val+1
}
