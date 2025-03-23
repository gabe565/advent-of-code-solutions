package day4

import (
	"bufio"
	"bytes"
	"io"
	"log/slog"
	"strconv"
)

type Bingo struct {
	Order     Order
	DrawIndex int
	Boards    []Board
}

func (bingo *Bingo) Write(p []byte) (int, error) {
	var n int
	for line := range bytes.SplitSeq(p, []byte{'\n'}) {
		if len(bingo.Order) == 0 {
			for v := range bytes.SplitSeq(line, []byte{','}) {
				parsed, err := strconv.Atoi(string(v))
				if err != nil {
					return n, err
				}
				bingo.Order = append(bingo.Order, parsed)
				n += len(v)
			}
		}
	}
	return n, nil
}

func (bingo *Bingo) ReadFrom(r io.Reader) (int64, error) {
	var n int64
	s := bufio.NewScanner(r)
	for i := 0; s.Scan(); i++ {
		switch i {
		case 0:
			written, err := bingo.Order.Write(s.Bytes())
			if err != nil {
				return n, err
			}
			n += int64(written)
		default:
			if len(s.Bytes()) == 0 {
				if len(bingo.Boards) > 0 {
					slog.Debug("Parsed board",
						"board", bingo.Boards[len(bingo.Boards)-1],
					)
				}
				bingo.Boards = append(bingo.Boards, Board{})
				continue
			}
			board := &bingo.Boards[len(bingo.Boards)-1]
			written, err := board.Write(s.Bytes())
			n += int64(written)
			if err != nil {
				return n, err
			}
		}
	}
	slog.Debug("Parsed board",
		"board", bingo.Boards[len(bingo.Boards)-1],
	)
	return n, s.Err()
}

func (bingo Bingo) String() string {
	var result string
	result += bingo.Order.String()
	result += "\n\n"
	for i, board := range bingo.Boards {
		result += board.String()
		if i < len(bingo.Boards)-1 {
			result += "\n\n"
		}
	}
	return result
}

func (bingo *Bingo) Draw() bool {
	if bingo.DrawIndex > len(bingo.Order)-1 {
		return false
	}
	drawn := bingo.Order[bingo.DrawIndex]
	bingo.DrawIndex++
	for i := range bingo.Boards {
		board := &bingo.Boards[i]
		for j := range *board {
			row := &(*board)[j]
			for k := range *row {
				cell := &(*row)[k]
				if cell.Value == drawn {
					cell.Drawn = true
				}
			}
		}
	}
	return true
}

func (bingo *Bingo) Winner() (int, *Board) {
	for i, board := range bingo.Boards {
		if board.Wins() {
			return i, &board
		}
	}
	return -1, nil
}
