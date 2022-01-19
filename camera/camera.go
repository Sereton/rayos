package camera

import (
	"bolijollo/rayos/canvas"
	"bolijollo/rayos/matrix"
	"bolijollo/rayos/primitives"
	"bolijollo/rayos/shapes"
	"bolijollo/rayos/world"
	"math"
)

type Camera struct {
	Transform     matrix.Matrix
	HSize         int
	VSize         int
	field_of_view float64
	pixel_size    float64
	half_width    float64
	half_height   float64
}

func CreateCamera(hsize, vsize int, field_of_view float64) Camera {
	var half_width, half_height, pixel_size float64
	half_view := math.Tan(field_of_view / 2)
	aspect := float64(hsize) / float64(vsize)
	if aspect >= 1 {
		half_width = half_view
		half_height = half_view / aspect
	} else {
		half_width = half_view * aspect
		half_height = half_view
	}
	pixel_size = (half_width * 2) / float64(hsize)
	return Camera{
		Transform:     matrix.Identity(4),
		HSize:         hsize,
		VSize:         vsize,
		field_of_view: field_of_view,
		pixel_size:    pixel_size,
		half_width:    half_width,
		half_height:   half_height,
	}
}

func (c *Camera) ray_for_pixel(px, py int) shapes.Ray {
	x_offset := (float64(px) + 0.5) * c.pixel_size
	y_offset := (float64(py) + 0.5) * c.pixel_size

	world_x := c.half_width - x_offset
	world_y := c.half_height - y_offset

	point := primitives.Point(world_x, world_y, -1)
	origin_point := primitives.Point(0, 0, 0)
	pixel := c.Transform.Inverse().VMulti(&point)
	origin := c.Transform.Inverse().VMulti(&origin_point)
	direction := pixel.Sub(origin).Normalize()

	return shapes.Ray{Origin: origin, Direction: direction}
}

func (c *Camera) Render(w *world.World, name string) {
	image := canvas.NewCanvas(c.HSize, c.VSize)
	for i := 0; i < c.HSize; i++ {
		for j := 0; j < c.VSize; j++ {
			ray := c.ray_for_pixel(i, j)
			color := world.Color_At(&ray, w)
			image.WritePixel(i, j, color)
		}

	}
	image.WriteToFile(name)
}
