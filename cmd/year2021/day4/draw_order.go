package day4

import (
	"bytes"
	"strconv"
)

type Order []int

func (order *Order) Write(p []byte) (int, error) {
	var n int
	for v := range bytes.SplitSeq(p, []byte{','}) {
		if len(v) == 0 {
			continue
		}
		parsed, err := strconv.Atoi(string(v))
		if err != nil {
			return n, err
		}
		*order = append(*order, parsed)
		n += len(v)
	}
	return n, nil
}

func (order Order) String() string {
	var result string
	for i, v := range order {
		result += strconv.Itoa(v)
		if i < len(order)-1 {
			result += ","
		}
	}
	return result
}
