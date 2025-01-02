package day6

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"slices"
	"strconv"

	"github.com/gabe565/advent-of-code-solutions/internal/util"
)

type Race struct {
	Time   []byte
	Record []byte
}

var ErrInvalidInput = errors.New("invalid input")

func (r *Race) Decode(rd io.Reader) error {
	scanner := bufio.NewScanner(rd)
	for i := 0; scanner.Scan(); i++ {
		switch i {
		case 0:
			r.Time = slices.Clone(bytes.TrimSpace(bytes.TrimPrefix(scanner.Bytes(), []byte("Time:"))))
		case 1:
			r.Record = slices.Clone(bytes.TrimSpace(bytes.TrimPrefix(scanner.Bytes(), []byte("Distance:"))))
		}
	}
	return scanner.Err()
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
