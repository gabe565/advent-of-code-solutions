package day4

import (
	"bytes"
	"strconv"
	"strings"
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
	var result strings.Builder
	for i, v := range order {
		result.WriteString(strconv.Itoa(v))
		if i < len(order)-1 {
			result.WriteString(",")
		}
	}
	return result.String()
}
