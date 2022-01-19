package scenes

import (
	"bolijollo/rayos/camera"
	"bolijollo/rayos/color"
	"bolijollo/rayos/lights"
	"bolijollo/rayos/materials"
	"bolijollo/rayos/matrix"
	"bolijollo/rayos/primitives"
	"bolijollo/rayos/shapes"
	"bolijollo/rayos/world"

	"math"
)

func DrawFirstWorld() {

	var height int = 40 * 10
	var width int = 80 * 10
	camara := camera.CreateCamera(width, height, math.Pi/3)

	camara.Transform = camera.View_transformation(primitives.Point(0, 1.5, -5),
		primitives.Point(0, 1, 0),
		primitives.Vector(0, -1, 0))

	// Create a world
	world := world.CreateWorld()

	// Add a point light source
	light := lights.CreatePointLight(primitives.Point(10, 10, -10), color.White)
	world.AddLight(light)

	// Add a floor
	floor := shapes.UniqueSphere()
	floor.T_Matrix = matrix.Scaling(10, 0.01, 10)
	floor.Material = materials.DefaultMaterial()
	floor.Material.Color = color.NewColor(1, 0.9, 0.9).Scale(255)
	floor.Material.Specular = 0
	world.AddObject(floor)

	// Add a left wall

	left_wall := shapes.UniqueSphere()
	left_wall.T_Matrix = left_wall.T_Matrix.MMulti(matrix.Translation(0, 0, 5)).
		MMulti(matrix.Rotation_Y(math.Pi / 4)).
		MMulti(matrix.Rotation_X(math.Pi / 2)).
		MMulti(matrix.Scaling(10, 0.01, 10))

	left_wall.Material = materials.DefaultMaterial()

	world.AddObject(left_wall)

	// Add a right wall
	right_wall := shapes.UniqueSphere()
	right_wall.T_Matrix = right_wall.T_Matrix.MMulti(matrix.Translation(0, 0, 5)).
		MMulti(matrix.Rotation_Y(-math.Pi / 4)).
		MMulti(matrix.Rotation_X(math.Pi / 2)).
		MMulti(matrix.Scaling(10, 0.01, 10))

	right_wall.Material = materials.DefaultMaterial()
	world.AddObject(right_wall)

	//Add a middle sphere

	middle_sphere := shapes.UniqueSphere()
	middle_sphere.T_Matrix = middle_sphere.T_Matrix.MMulti(matrix.Translation(0.5, 1, 0.5))
	middle_sphere.Material = materials.DefaultMaterial()
	middle_sphere.Material.Color = color.NewColor(0.1, 1, 0.5).Scale(255)
	middle_sphere.Material.Diffuse = 0.7
	middle_sphere.Material.Specular = 0.3
	world.AddObject(middle_sphere)

	// Add a right sphere
	right_sphere := shapes.UniqueSphere()
	right_sphere.T_Matrix = right_sphere.T_Matrix.MMulti(matrix.Translation(-1, 0.5, -0.5)).
		MMulti(matrix.Scaling(0.5, 0.5, 0.5))
	right_sphere.Material = materials.DefaultMaterial()
	right_sphere.Material.Color = color.NewColor(0.5, 1, 0.1).Scale(255)
	right_sphere.Material.Diffuse = 0.7
	right_sphere.Material.Specular = 0.3

	world.AddObject(right_sphere)

	// Add a left sphere
	left_sphere := shapes.UniqueSphere()
	left_sphere.T_Matrix = left_sphere.T_Matrix.MMulti(matrix.Translation(1.5, 0.33, -0.75)).
		MMulti(matrix.Scaling(0.33, 0.33, 0.33))
	left_sphere.Material = materials.DefaultMaterial()
	left_sphere.Material.Color = color.NewColor(1, 0.8, 0.1).Scale(255)
	left_sphere.Material.Diffuse = 0.7
	left_sphere.Material.Specular = 0.3

	world.AddObject(left_sphere)

	camara.Render(world, "first_world.ppm")
}
