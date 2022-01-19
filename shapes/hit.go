package shapes

import "math"

func Hit(intersections []Intersection) Intersection {

	var hit Intersection
	var nearest_intersection Intersection
	var nearest_distance float64 = math.MaxFloat64
	for _, intersection := range intersections {
		if intersection.T < nearest_distance && intersection.T > 0 {
			nearest_intersection = intersection
			nearest_distance = intersection.T
		}
	}
	if nearest_distance < math.MaxFloat64 {
		hit = nearest_intersection
	}
	return hit
}
