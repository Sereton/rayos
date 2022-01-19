package world

import (
	"bolijollo/rayos/primitives"
	"bolijollo/rayos/shapes"
)

func Is_Shadowed(w World, p primitives.Tuple) bool {

	for _, l := range w.Lights {
		light_distance_vec := l.Position.Sub(p)
		distance := light_distance_vec.Length()
		light_distance_normalized_vec := light_distance_vec.Normalize()

		shadow_ray := shapes.Ray{
			Origin:    p,
			Direction: light_distance_normalized_vec,
		}

		intersections := w.Intersect(&shadow_ray)
		if intersections != nil {
			hit := shapes.Hit(intersections).T
			if hit < distance && hit > 0.0000000001 {
				return true
			}
		}
	}

	return false
}
