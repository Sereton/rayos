package scenes

import (
	"bolijollo/rayos/canvas"
	"bolijollo/rayos/color"
	"bolijollo/rayos/lights"

	p "bolijollo/rayos/primitives"
	"bolijollo/rayos/shapes"
)

func DrawLightenedSphere() {
	var center_canvas p.Tuple = p.Point(0, 0, 3)
	var camera_origin p.Tuple = p.Point(0, 0, -7)
	sphere := shapes.UniqueSphere()
	sphere.Material.Color = color.NewColor(255, 51, 255)
	light := lights.CreatePointLight(p.Point(-10, 10, -10), color.White)
	var height int = 350
	var width int = 350
	var max_height int = height / 2
	var max_width int = width / 2
	pixels_per_meter := 50.0
	canvas := canvas.NewCanvas(width, height)
	for i := 0; i < int(height); i++ {
		for j := 0; j < int(width); j++ {
			var ray_origin p.Tuple = p.Vector(float64(i-max_width)/pixels_per_meter,
				float64(j-max_height)/pixels_per_meter, 0)
			var canvas_point = ray_origin.Add(center_canvas)

			var ray_direction = canvas_point.Sub(camera_origin).Normalize()
			var ray shapes.Ray = shapes.Ray{Origin: camera_origin, Direction: ray_direction}
			var intersections = ray.IntersectSphere(sphere)
			hit := shapes.Hit(sphere, intersections)
			if hit.Object != nil {
				// fmt.Println("HIT", i, j, shapes.Hit(sphere, intersections).T)
				point := ray.Position_at(hit.T)
				normal := sphere.NormalAt(&point)
				eye := ray.Direction.Minus()
				color := lights.Lighting(sphere.Material, light, point, eye, normal)
				canvas.WritePixel(i, j, color)

			}
		}
	}
	canvas.WriteToFile("lightened_sphere.ppm")
}
