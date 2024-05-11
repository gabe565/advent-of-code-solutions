package day9

import (
	"bufio"
	"io"

	"github.com/gabe565/advent-of-code-2023/internal/util"
)

type Report struct {
	History []History
}

func (r *Report) Decode(rd io.Reader) error {
	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		var history History
		if err := history.UnmarshalText(scanner.Bytes()); err != nil {
			return err
		}
		r.History = append(r.History, history)
	}
	return scanner.Err()
}

func (r *Report) Predict(mode PredictMode) int {
	vals := make([]int, 0, len(r.History))
	for _, history := range r.History {
		vals = append(vals, history.Predict(mode))
	}
	return util.Sum(vals)
}

type History struct {
	Values []int
}

func (h *History) UnmarshalText(text []byte) error {
	var err error
	h.Values, err = util.StringToIntSlice(string(text), " ")
	return err
}

type PredictMode uint8

const (
	PredictFuture PredictMode = iota
	PredictPast
)

func (h *History) Predict(mode PredictMode) int {
	lines := [][]int{h.Values}
	for util.Sum(lines[len(lines)-1]) != 0 {
		prevLine := lines[len(lines)-1]
		thisLine := make([]int, 0, len(prevLine))
		for i := range len(prevLine) - 1 {
			thisLine = append(thisLine, prevLine[i+1]-prevLine[i])
		}
		lines = append(lines, thisLine)
	}

	lines[len(lines)-1] = append(lines[len(lines)-1], 0)
	for i := len(lines) - 1; i > 0; i-- {
		prevLine := lines[i-1]
		thisLine := lines[i]
		switch mode {
		case PredictFuture:
			prevLine = append(prevLine, thisLine[len(thisLine)-1]+prevLine[len(prevLine)-1])
		case PredictPast:
			prevLine = append([]int{prevLine[0] - thisLine[0]}, prevLine...)
		}
		lines[i-1] = prevLine
	}
	switch mode {
	case PredictFuture:
		return lines[0][len(lines[0])-1]
	case PredictPast:
		return lines[0][0]
	}
	return 0
}
