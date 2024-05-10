package day3

import (
	"bufio"
	"io"
	"regexp"
	"slices"
	"strconv"
)

var splitRe = regexp.MustCompile("[^0-9]")

type Schematic struct {
	Grid    [][]byte
	Numbers []Number
}

func (s *Schematic) Decode(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for x := 0; scanner.Scan(); x++ {
		s.Grid = append(s.Grid, slices.Clone(scanner.Bytes()))
		split := splitRe.Split(scanner.Text(), -1)
		var y int
		for _, val := range split {
			if len(val) != 0 {
				parsed, err := strconv.Atoi(val)
				if err != nil {
					continue
				}
				s.Numbers = append(s.Numbers, Number{
					Value: parsed,
					X:     x,
					Y:     y,
					Len:   len(val),
				})
			}
			y += 1 + len(val)
		}
	}
	return scanner.Err()
}

func (s *Schematic) Ratios() (int, int) {
	var part1, part2 int

	for x, line := range s.Grid {
		for y, char := range line {
			if char != '.' && char < '0' || char > '9' {
				computeRatio := char == '*'
				localMatches := make([]int, 0, 2)
				adjacent := [][]byte{
					s.Grid[x-1][y-1 : y+2],
					s.Grid[x][y-1 : y+2],
					s.Grid[x+1][y-1 : y+2],
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
							for _, n := range s.Numbers {
								if realX == n.X && realY >= n.Y && realY <= n.Y+n.Len {
									part1 += n.Value
									skip = n.Y - realY + n.Len - 1
									if computeRatio {
										localMatches = append(localMatches, n.Value)
									}
								}
							}
							if len(localMatches) == 2 {
								part2 += localMatches[0] * localMatches[1]
							}
						}
					}
				}
			}
		}
	}

	return part1, part2
}

type Number struct {
	Value int
	X, Y  int
	Len   int
}
