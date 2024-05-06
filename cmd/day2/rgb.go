package day2

import (
	"bytes"
	"fmt"
)

type RGB struct {
	R, G, B int
}

func (r *RGB) UnmarshalText(text []byte) error {
	for _, colors := range bytes.Split(text, []byte(",")) {
		var color string
		var number int
		if _, err := fmt.Fscanf(bytes.NewReader(colors), "%d %s", &number, &color); err != nil {
			return err
		}

		switch color {
		case "red":
			r.R = number
		case "green":
			r.G = number
		case "blue":
			r.B = number
		default:
			return fmt.Errorf("invalid color: %s", color)
		}
	}
	return nil
}

func (r *RGB) Valid() bool {
	return r.R <= RedCubes && r.G <= GreenCubes && r.B <= BlueCubes
}
