package day8

import (
	"bufio"
	"image"
	"io"
	"slices"
	"strings"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() day.Day[*Map, int] {
	return day.Day[*Map, int]{
		Year: 2024,
		Day:  8,
		Parse: func(r io.Reader) (*Map, error) {
			var m Map
			scanner := bufio.NewScanner(r)
			for y := 0; scanner.Scan(); y++ {
				m.Rect.Max.X = len(scanner.Bytes())
				m.Rect.Max.Y = y + 1

				for x, b := range scanner.Bytes() {
					if b != Empty {
						m.Antennas = append(m.Antennas, Antenna{
							Point:  image.Pt(x, y),
							Letter: b,
						})
					}
				}
			}
			return &m, scanner.Err()
		},
		Part1: func(input *Map) (int, error) {
			input.Harmonics = false
			return len(input.GetAllAntinodes()), nil
		},
		Part2: func(input *Map) (int, error) {
			input.Harmonics = true
			return len(input.GetAllAntinodes()), nil
		},
	}
}

const (
	Empty    byte = '.'
	Antinode byte = '#'
)

type Map struct {
	Rect      image.Rectangle
	Antennas  []Antenna
	Harmonics bool
}

type Antenna struct {
	Point  image.Point
	Letter byte
}

func (m *Map) AntennaCodes() []byte {
	codes := make([]byte, 0, len(m.Antennas))
	for _, a := range m.Antennas {
		if !slices.Contains(codes, a.Letter) {
			codes = append(codes, a.Letter)
		}
	}
	return codes
}

func (m *Map) AntennasByCode(code byte) []Antenna {
	antennas := make([]Antenna, 0, len(m.Antennas))
	for _, a := range m.Antennas {
		if a.Letter == code {
			antennas = append(antennas, a)
		}
	}
	return antennas
}

func (m *Map) GetAntinodes(code byte) []image.Point {
	antennas := m.AntennasByCode(code)
	antinodes := make([]image.Point, 0, len(antennas)*2)
	for i, a := range antennas {
		for j, b := range antennas {
			if i == j {
				continue
			}
			diff := a.Point.Sub(b.Point)
			if m.Harmonics {
				pt := b.Point
				for pt.In(m.Rect) {
					antinodes = append(antinodes, pt)
					pt = pt.Sub(diff)
				}
			} else if pt := b.Point.Sub(diff); pt.In(m.Rect) {
				antinodes = append(antinodes, pt)
			}
		}
	}
	return antinodes
}

func (m *Map) GetAllAntinodes() []image.Point {
	antinodes := make([]image.Point, 0, len(m.Antennas))
	for _, code := range m.AntennaCodes() {
		for _, antinode := range m.GetAntinodes(code) {
			if !slices.Contains(antinodes, antinode) {
				antinodes = append(antinodes, antinode)
			}
		}
	}
	return antinodes
}

func (m *Map) String() string {
	var s strings.Builder
	s.Grow(m.Rect.Dx() * m.Rect.Dy())

	antinodes := m.GetAllAntinodes()

	for y := range m.Rect.Max.Y {
		for x := range m.Rect.Max.X {
			b := Empty
			for _, a := range m.Antennas {
				if a.Point.X == x && a.Point.Y == y {
					b = a.Letter
					break
				}
				for _, antinode := range antinodes {
					if antinode.X == x && antinode.Y == y {
						if b == Empty {
							b = Antinode
							break
						}
					}
				}
			}
			s.WriteByte(b)
		}
		s.WriteByte('\n')
	}
	return s.String()
}
