package day6

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"

	"github.com/gabe565/advent-of-code-2023/internal/util"
)

type Race struct {
	Time   []byte
	Record []byte
}

var ErrInvalidInput = errors.New("invalid input")

func (r *Race) UnmarshalText(text []byte) error {
	timeLine, recordLine, ok := bytes.Cut(text, []byte("\n"))
	if !ok {
		return fmt.Errorf("%w: %s", ErrInvalidInput, string(text))
	}

	r.Time = bytes.TrimSpace(bytes.TrimPrefix(timeLine, []byte("Time:")))
	r.Record = bytes.TrimSpace(bytes.TrimPrefix(recordLine, []byte("Distance:")))
	return nil
}

func (r *Race) Part1() (int, error) {
	totalCount := 1
	times, err := util.StringToIntSlice(string(r.Time), " ")
	if err != nil {
		return 0, err
	}
	records, err := util.StringToIntSlice(string(r.Record), " ")
	if err != nil {
		return 0, err
	}

	for i, time := range times {
		totalCount *= countWins(time, records[i])
	}

	return totalCount, nil
}

func (r *Race) Part2() (int, error) {
	var timeRaw, recordRaw []byte
	for _, v := range r.Time {
		if v != ' ' {
			timeRaw = append(timeRaw, v)
		}
	}
	for _, v := range r.Record {
		if v != ' ' {
			recordRaw = append(recordRaw, v)
		}
	}

	time, err := strconv.Atoi(string(timeRaw))
	if err != nil {
		return 0, err
	}
	record, err := strconv.Atoi(string(recordRaw))
	if err != nil {
		return 0, err
	}

	return countWins(time, record), nil
}

func countWins(time, record int) int {
	var count int
	for speed := range time {
		holdDuration := speed
		driveTime := time - holdDuration
		distance := speed * driveTime
		if distance > record {
			count++
		}
	}
	return count
}
