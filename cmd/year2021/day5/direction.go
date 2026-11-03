package day5

type Direction uint8

const (
	Unknown Direction = iota
	Point
	Vertical
	Horizontal
	Diagonal
)
