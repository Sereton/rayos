package shapes

type Shape interface {
	Intersect(ray *Ray) []float64
}
