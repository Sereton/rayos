package lights

import (
	"bolijollo/rayos/color"
	"bolijollo/rayos/materials"
	"bolijollo/rayos/primitives"
	"math"
)

type PointLight struct {
	Position  primitives.Tuple
	Intensity color.Color
}

func CreatePointLight(position primitives.Tuple, color color.Color) PointLight {
	return PointLight{Position: position, Intensity: color}
}

func Lighting(material materials.Material, light *PointLight, Point primitives.Tuple, EyeV primitives.Tuple, NormalV primitives.Tuple) color.Color {
	effectiveColor := color.Hadamard(&material.Color, &light.Intensity)

	// find direction to light source
	lightV := light.Position.Sub(Point).Normalize()

	// compute the ambient contribution
	diffuse := color.Black
	specular := color.Black
	ambient := effectiveColor.Scale(material.Ambient)

	// Calculate light_dot_normal
	light_dot_normal := lightV.Dot(NormalV)

	if light_dot_normal < 0 {
		//nothing so far because diffuse and specular are already black
	} else {
		// diffuse contribution
		diffuse = effectiveColor.Scale(material.Diffuse * light_dot_normal)

		// reflect_vector = 2 * light_dot_normal * normal - light_direction, the function
		// is implemented in the primitives package
		reflect_v := lightV.Minus().Reflect(NormalV)

		// reflect_dot_eye = reflect_vector.dot(eye)
		reflect_dot_eye := reflect_v.Dot(EyeV)

		if reflect_dot_eye <= 0 {
			specular = color.Black
		} else {
			// specular contribution
			factor := math.Pow(reflect_dot_eye, material.Shininess)
			specular = light.Intensity.Scale(material.Specular * factor)
		}

		//
	}
	return ambient.Add(diffuse).Add(specular)
}
