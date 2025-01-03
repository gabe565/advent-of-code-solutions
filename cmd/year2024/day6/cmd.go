package day6

import (
	"bytes"
	"image"
	"io"
	"slices"
	"strings"
	"time"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
)

func New() day.Day[*Map, int] {
	return day.Day[*Map, int]{
		Date: time.Date(2024, 12, 6, 0, 0, 0, 0, time.Local),
		Parse: func(r io.Reader) (*Map, error) {
			b, err := io.ReadAll(r)
			if err != nil {
				return nil, err
			}
			b = bytes.TrimSpace(b)

			var m Map
			m.Pix = bytes.Split(b, []byte("\n"))
			m.Rect.Max = image.Pt(len(m.Pix[0]), len(m.Pix))

			for y, line := range m.Pix {
				for x, b := range line {
					switch b {
					case Obstacle, SpecialObstacle:
					case Empty:
						m.Pix[y][x] = 0
					case GuardNorth, GuardEast, GuardSouth, GuardWest:
						m.Pix[y][x] = 0
						m.Guard.Pos = image.Pt(x, y)
						switch b {
						case GuardNorth:
							m.Guard.Direction = DirectionNorth
						case GuardEast:
							m.Guard.Direction = DirectionEast
						case GuardSouth:
							m.Guard.Direction = DirectionSouth
						case GuardWest:
							m.Guard.Direction = DirectionWest
						}
					}
				}
			}
			return &m, nil
		},
		Part1: func(m *Map) (int, error) {
			m.RunForward()
			return m.CellsVisited(), nil
		},
		Part2: func(m *Map) (int, error) {
			var result int
			for {
				if !m.Guard.Pos.In(m.Rect) {
					break
				}
				nextPos := m.Guard.NextPos()
				if !nextPos.In(m.Rect) {
					break
				}
				if m.GetCell(nextPos) == 0 {
					m.SaveState()
					m.SetCell(nextPos, SpecialObstacle)
					if !m.RunForward() {
						result++
					}
					m.RestoreState()
				}
				m.Step()
			}
			return result, nil
		},
	}
}

type Map struct {
	Pix   [][]byte
	Rect  image.Rectangle
	Guard Guard

	prevPix   [][]byte
	prevGuard Guard
}

func (m *Map) RunForward() bool {
	for {
		m.Step()
		switch {
		case !m.Guard.Pos.In(m.Rect):
			return true
		case m.Visits(m.Guard.Pos) > 10:
			return false
		}
	}
}

func (m *Map) SaveState() {
	m.prevPix = make([][]byte, 0, len(m.Pix))
	for _, line := range m.Pix {
		m.prevPix = append(m.prevPix, slices.Clone(line))
	}
	m.prevGuard = m.Guard
}

func (m *Map) RestoreState() {
	if m.prevPix == nil {
		panic("no state was saved")
	}
	m.Pix, m.prevPix = m.prevPix, nil
	m.Guard = m.prevGuard
}

func (m *Map) GetCell(pt image.Point) byte {
	return m.Pix[pt.Y][pt.X]
}

func (m *Map) SetCell(pt image.Point, value byte) {
	m.Pix[pt.Y][pt.X] = value
}

func (m *Map) Step() {
	m.SetCell(m.Guard.Pos, m.GetCell(m.Guard.Pos)+1)
	if m.IsObstacle(m.Guard.NextPos()) {
		m.Guard.Turn()
	} else {
		m.Guard.Move()
	}
}

func (m *Map) CellsVisited() int {
	var count int
	for _, line := range m.Pix {
		for _, b := range line {
			switch b {
			case 0, Obstacle, SpecialObstacle:
			default:
				count++
			}
		}
	}
	return count
}

func (m *Map) Visits(pt image.Point) int {
	if !pt.In(m.Rect) {
		return 0
	}
	cell := m.GetCell(pt)
	switch cell {
	case Obstacle, SpecialObstacle:
		panic("guard can't visit an obstacle")
	default:
		return int(cell)
	}
}

func (m *Map) IsObstacle(pt image.Point) bool {
	var pix byte
	if pt.In(m.Rect) {
		pix = m.GetCell(pt)
	}
	return pix == Obstacle || pix == SpecialObstacle
}

func (m *Map) String() string {
	var s strings.Builder
	s.Grow(m.Rect.Dx() * m.Rect.Dy())
	for y, line := range m.Pix {
		for x, b := range line {
			switch {
			case m.Guard.Pos.X == x && m.Guard.Pos.Y == y:
				s.WriteString(m.Guard.Direction.String())
			case b == Obstacle, b == SpecialObstacle:
				s.WriteByte(b)
			case b == 0:
				s.WriteByte(Empty)
			default:
				s.WriteByte(Visited)
			}
		}
		s.WriteByte('\n')
	}
	return s.String()
}

const (
	Empty           = '.'
	Visited         = 'X'
	Obstacle        = '#'
	SpecialObstacle = 'O'
	GuardNorth      = '^'
	GuardEast       = '>'
	GuardSouth      = 'v'
	GuardWest       = '<'
)

type Direction byte

const (
	DirectionNorth Direction = iota
	DirectionEast
	DirectionSouth
	DirectionWest
)

func (d Direction) String() string {
	switch d {
	case DirectionNorth:
		return string(GuardNorth)
	case DirectionEast:
		return string(GuardEast)
	case DirectionSouth:
		return string(GuardSouth)
	case DirectionWest:
		return string(GuardWest)
	default:
		panic("invalid direction")
	}
}

type Guard struct {
	Pos       image.Point
	Direction Direction
}

func (g *Guard) NextPos() image.Point {
	switch g.Direction {
	case DirectionNorth:
		return g.Pos.Sub(image.Pt(0, 1))
	case DirectionEast:
		return g.Pos.Add(image.Pt(1, 0))
	case DirectionSouth:
		return g.Pos.Add(image.Pt(0, 1))
	case DirectionWest:
		return g.Pos.Sub(image.Pt(1, 0))
	}
	panic("invalid direction")
}

func (g *Guard) Move() *Guard {
	g.Pos = g.NextPos()
	return g
}

func (g *Guard) Turn() *Guard {
	g.Direction++
	g.Direction %= DirectionWest + 1
	return g
}
