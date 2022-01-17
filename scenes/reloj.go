package scenes

import (
	"bolijollo/rayos/canvas"
	"bolijollo/rayos/color"
	"bolijollo/rayos/matrix"
	p "bolijollo/rayos/primitives"
	"math"
)

var height = 500
var width = 500
var origin p.Tuple = p.Point(float64(height/2), float64(width/2), 0)

var azul = color.NewColor(255, 0, 255)

var rotation_Matrix = matrix.Rotation_Z(math.Pi / 50)
var vec_to_rotate = p.Vector(0, float64(height*3)/8, 0)

func DrawClock() {
	canvas := canvas.NewCanvas(width, height)

	for counter := 0; counter < 100; counter++ {
		hour_mark := origin.Add(&vec_to_rotate)

		x := int(hour_mark.X)
		y := int(hour_mark.Y)
		canvas.WritePixel(x, y, azul)

		vec_to_rotate = rotation_Matrix.VMulti(&vec_to_rotate)

	}
	canvas.WriteToFile("clock.ppm")
}
