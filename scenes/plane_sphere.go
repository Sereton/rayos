package scenes

import (
	"bolijollo/rayos/canvas"
	"bolijollo/rayos/color"
	"bolijollo/rayos/matrix"
	p "bolijollo/rayos/primitives"
	"bolijollo/rayos/shapes"
)

func DrawSphere() {
	var rojo = color.NewColor(255, 0, 0)
	// var azul = color.NewColor(0, 0, 255)
	var center_canvas p.Tuple = p.Point(0, 0, 10)
	var camera_origin p.Tuple = p.Point(0, 0, -5)

	// 1 meter = 20 pixels
	var height int = 350
	var width int = 350
	var max_height int = height / 2
	var max_width int = width / 2
	pixels_per_meter := 50.0
	sphere := shapes.UniqueSphere()
	sphere = sphere.ChangeTransformation(matrix.Scaling(0.5, 1, 1)).ChangeTransformation(matrix.Shearing(1, 0, 0, 0, 0, 0)).ChangeTransformation(matrix.Translation(0, 0, -1))
	canvas := canvas.NewCanvas(width, height)
	for i := 0; i < int(height); i++ {
		for j := 0; j < int(width); j++ {
			var ray_origin p.Tuple = p.Vector(float64(i-max_width)/pixels_per_meter,
				float64(j-max_height)/pixels_per_meter, 0)
			var canvas_point = ray_origin.Add(&center_canvas)

			var ray_direction = canvas_point.Sub(&camera_origin).Normalize()
			var ray shapes.Ray = shapes.Ray{Origin: camera_origin, Direction: ray_direction}
			var intersections = ray.IntersectSphere(sphere)

			if shapes.Hit(sphere, intersections).Object != nil {
				// fmt.Println("HIT", i, j, shapes.Hit(sphere, intersections).T)

				canvas.WritePixel(i, j, rojo)

			}
		}
	}
	canvas.WriteToFile("flat_sphere.ppm")
}
