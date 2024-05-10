package day8

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/gabe565/advent-of-code-2023/internal/util"
)

type Network struct {
	Directions []int
	Map        map[string][2]string
}

func (n *Network) Decode(r io.Reader) error {
	n.Map = make(map[string][2]string)
	mapRe := regexp.MustCompile(`^(?P<name>.*) = \((?P<left>.*), (?P<right>.*)\)$`)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			continue
		}
		if len(n.Directions) == 0 {
			for _, b := range scanner.Bytes() {
				var dir int
				switch b {
				case 'L':
					dir = 0
				case 'R':
					dir = 1
				default:
					return fmt.Errorf("invalid direction: %s", string(b))
				}
				n.Directions = append(n.Directions, dir)
			}
		} else {
			matches := mapRe.FindSubmatch(scanner.Bytes())
			if len(matches) == 0 {
				return fmt.Errorf("no match found for %s", scanner.Text())
			}

			var name, left, right string
			for i, field := range mapRe.SubexpNames() {
				switch field {
				case "name":
					name = string(matches[i])
				case "left":
					left = string(matches[i])
				case "right":
					right = string(matches[i])
				}
			}
			n.Map[name] = [2]string{left, right}
		}
	}
	return scanner.Err()
}

func (n *Network) Steps(at string, dstEndsWith bool) (int, error) {
	var steps int
	var dirIdx int
	for {
		if dstEndsWith {
			if strings.HasSuffix(at, "Z") {
				return steps, nil
			}
		} else {
			if at == "ZZZ" {
				return steps, nil
			}
		}

		dir := n.Directions[dirIdx]
		if current, ok := n.Map[at]; ok {
			at = current[dir]
		} else {
			return 0, fmt.Errorf("no map entry found for %s", at)
		}
		dirIdx++
		dirIdx %= len(n.Directions)
		steps++
	}
}

func (n *Network) GhostSteps() (int, error) {
	var steps []int
	for name := range n.Map {
		if strings.HasSuffix(name, "A") {
			n, err := n.Steps(name, true)
			if err != nil {
				return 0, err
			}
			steps = append(steps, n)
		}
	}
	return util.LCM(steps...)
}
