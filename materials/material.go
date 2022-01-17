package materials

import (
	"bolijollo/rayos/color"
)

type Material struct {
	Color                                 color.Color
	Ambient, Diffuse, Specular, Shininess float64
}

func DefaultMaterial() Material {
	return Material{
		Color:     color.White,
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0}
}
