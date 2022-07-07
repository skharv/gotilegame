package geom

import "math"

type Vector2[T int | float64] struct {
	X, Y T
}

func (v *Vector2[T]) Lerp(target Vector2[T], step float64) {
	x := float64(v.X) + float64(target.X-v.X)*step
	y := float64(v.Y) + float64(target.Y-v.Y)*step

	v.X = T(x)
	v.Y = T(y)
}

func (v *Vector2[T]) DistanceTo(target Vector2[T]) float64 {
	return math.Sqrt(math.Pow(float64(v.X-target.X), 2) + math.Pow(float64(v.Y-target.Y), 2))
}
