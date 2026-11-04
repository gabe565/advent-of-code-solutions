package day7

import (
	"io"
	"math"
	"slices"

	"github.com/gabe565/advent-of-code-solutions/internal/day"
	"github.com/gabe565/advent-of-code-solutions/internal/util"
)

func New() *day.Day[[]int, int] {
	return &day.Day[[]int, int]{
		Year: 2021,
		Day:  7,
		Parse: func(r io.Reader) ([]int, error) {
			data, err := io.ReadAll(r)
			if err != nil {
				return nil, err
			}

			input, err := util.StringToIntSlice(string(data), ",")
			if err != nil {
				return nil, err
			}

			slices.Sort(input)

			return input, nil
		},
		Part1: func(input []int) (int, error) {
			median := input[len(input)/2]

			var cost int
			for _, v := range input {
				cost += util.Abs(v - median)
			}

			return cost, nil
		},
		Part2: func(input []int) (int, error) {
			mean := float64(util.Sum(input)) / float64(len(input))
			floorMean := int(math.Floor(mean))
			ceilMean := int(math.Ceil(mean))

			var floorCost, ceilCost int
			for _, v := range input {
				floorDist := util.Abs(v - floorMean)
				ceilDist := util.Abs(v - ceilMean)
				floorCost += floorDist * (floorDist + 1) / 2
				ceilCost += ceilDist * (ceilDist + 1) / 2
			}

			return min(floorCost, ceilCost), nil
		},
	}
}
