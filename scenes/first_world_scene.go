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

	var height int = 40 * 12
	var width int = 80 * 12
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
	floor := shapes.UniquePlane()
	floor.TMatrix = matrix.Scaling(10, 0.01, 10)
	floor.Material = materials.DefaultMaterial()
	floor.Material.Color = color.NewColor(1, 1, 1).Scale(255)
	floor.Material.Specular = 1
	floor.Material.Diffuse = 1
	world.AddObject(floor)

	// Add a left wall

	left_wall := shapes.UniquePlane()
	left_wall.TMatrix = left_wall.TMatrix.MMulti(matrix.Translation(0, 0, 5)).
		MMulti(matrix.Rotation_Y(math.Pi / 4)).
		MMulti(matrix.Rotation_X(math.Pi / 2)).
		MMulti(matrix.Scaling(10, 0.01, 10))

	left_wall.Material = materials.DefaultMaterial()

	world.AddObject(left_wall)

	// Add a right wall
	right_wall := shapes.UniquePlane()
	right_wall.TMatrix = right_wall.TMatrix.MMulti(matrix.Translation(0, 0, 5)).
		MMulti(matrix.Rotation_Y(-math.Pi / 4)).
		MMulti(matrix.Rotation_X(math.Pi / 2)).
		MMulti(matrix.Scaling(10, 0.01, 10))

	right_wall.Material = materials.DefaultMaterial()
	world.AddObject(right_wall)

	//Add a middle sphere

	middle_sphere := shapes.UniqueSphere()
	middle_sphere.T_Matrix = middle_sphere.T_Matrix.MMulti(matrix.Translation(0.5, 1.5, 0.5))
	middle_sphere.Material = materials.DefaultMaterial()
	middle_sphere.Material.Color = color.NewColor(0, 84, 119)
	middle_sphere.Material.Diffuse = 0.7
	middle_sphere.Material.Specular = 0.8
	world.AddObject(middle_sphere)

	// Add a right sphere
	right_sphere := shapes.UniqueSphere()
	right_sphere.T_Matrix = right_sphere.T_Matrix.MMulti(matrix.Translation(-1.3, 1.2, -0.5)).
		MMulti(matrix.Scaling(0.2, 0.2, 0.2))
	right_sphere.Material = materials.DefaultMaterial()
	right_sphere.Material.Color = color.NewColor(216, 214, 203)
	right_sphere.Material.Diffuse = 0.7
	right_sphere.Material.Specular = 0.3

	world.AddObject(right_sphere)

	// Add a left sphere
	left_sphere := shapes.UniqueSphere()
	left_sphere.T_Matrix = left_sphere.T_Matrix.MMulti(matrix.Translation(2, 1.33, -0.75)).
		MMulti(matrix.Scaling(0.33, 0.33, 0.33))
	left_sphere.Material = materials.DefaultMaterial()
	left_sphere.Material.Color = color.NewColor(216, 214, 203)
	left_sphere.Material.Diffuse = 0.5
	left_sphere.Material.Specular = 0.3

	world.AddObject(left_sphere)

	camara.Render(world, "first_world.ppm")
}
