package algorithm

import (
	"math"

	"golang.org/x/exp/constraints"
)

type point2D[T constraints.Integer | constraints.Float] struct {
	X, Y T
}

// Distance calculate distance between two points
func (p *point2D[T]) Distance(x, y T) float64 {
	distanceX := float64(x) - float64(p.X)
	distanceY := float64(y) - float64(p.Y)

	pow2 := math.Pow(distanceX, 2) + math.Pow(distanceY, 2)

	return math.Sqrt(pow2)
}

type pointValue2D[T constraints.Integer | constraints.Float, V any] struct {
	point2D[T]
	Value V
}
