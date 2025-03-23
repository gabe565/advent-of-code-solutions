package day4

import (
	"bytes"
	"strconv"
)

type Board [][]Cell

func (board *Board) Write(p []byte) (int, error) {
	var n int
	for line := range bytes.SplitSeq(p, []byte{'\n'}) {
		split := bytes.Split(line, []byte{' '})
		row := make([]Cell, 0, len(split))
		for _, v := range split {
			if len(v) == 0 {
				continue
			}
			parsed, err := strconv.Atoi(string(v))
			if err != nil {
				return n, err
			}
			row = append(row, Cell{Value: parsed})
			n += len(v)
		}
		*board = append(*board, row)
	}
	return n, nil
}

func (board Board) String() string {
	var result string
	for i, row := range board {
		for _, cell := range row {
			result += cell.String()
		}
		if i < len(board)-1 {
			result += "\n"
		}
	}
	return result
}

func (board Board) Wins() bool {
	// row
	for _, row := range board {
		var valid bool
		for _, v := range row {
			valid = v.Drawn
			if !valid {
				break
			}
		}
		if valid {
			return true
		}
	}
	// column
	for i := range board[0] {
		var valid bool
		for j := range board {
			v := board[j][i]
			valid = v.Drawn
			if !valid {
				break
			}
		}
		if valid {
			return true
		}
	}
	return false
}

func (board Board) Sum() int {
	var total int
	for _, row := range board {
		for _, v := range row {
			if !v.Drawn {
				total += v.Value
			}
		}
	}
	return total
}
