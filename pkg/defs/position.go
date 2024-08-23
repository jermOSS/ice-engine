package defs

import (
	"math"
)

type Position struct {
	X int64
	Y int64
}

// Position constructor
func NewPosition(x int64, y int64) Position {
	return Position{
		X: x,
		Y: y,
	}
}

// Top-left to bottom-right diagonal
func (self Position) GetI() int64 {
	return self.X + self.Y
}

// Bottom-left to top-right diagonal
func (self Position) GetJ() int64 {
	return self.Y - self.X
}

func EuclideanDistance(p1 Position, p2 Position) float64 {
	return math.Sqrt(float64((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.X-p2.X)))
}

// https://en.wikipedia.org/wiki/Taxicab_geometry
func ManhattanDistance(p1 Position, p2 Position) float64 {
	return math.Abs(float64(p1.X-p2.X)) + math.Abs(float64(p1.Y-p2.Y))
}

// https://en.wikipedia.org/wiki/Chebyshev_distance
func ChebyshevDistance(p1 Position, p2 Position) float64 {
	return math.Max(math.Abs(float64(p1.X-p2.X)), math.Abs(float64(p1.Y-p2.Y)))
}
