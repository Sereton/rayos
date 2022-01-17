package shapes

type Intersection struct {
	T      float64
	Object Shape
}

func intersections(values ...Intersection) []Intersection {
	lista := make([]Intersection, len(values))
	for i, v := range values {
		lista[i] = v
	}
	return lista
}
