package canvas

import (
	"fmt"
	"rayos/color"
)

type Canvas struct {
	Width  int
	Height int
	Pixels [][]color.Color
}

func (c *Canvas) GetPixel(x, y int) color.Color {
	return c.Pixels[x][y]
}

func (c *Canvas) SetPixel(x, y int, color color.Color) {
	c.Pixels[x][y] = color
}

func NewCanvas(width, height int) Canvas {
	canvas := Canvas{Width: width, Height: height}
	canvas.Pixels = make([][]color.Color, width)
	for i := range canvas.Pixels {
		canvas.Pixels[i] = make([]color.Color, height)
	}
	return canvas
}

func (c *Canvas) WritePixel(x, y int, color color.Color) {
	c.Pixels[x][y] = color
}

func (c *Canvas) WritePPM() string {
	var ppm string
	ppm += fmt.Sprintf("P3\n%d %d\n%d\n", c.Width, c.Height, 255)
	for y := c.Height - 1; y >= 0; y-- {
		for x := 0; x < c.Width; x++ {
			ppm += fmt.Sprintf("%d %d %d ", int(c.Pixels[x][y].X), int(c.Pixels[x][y].Y),
				int(c.Pixels[x][y].Z))
		}
		ppm += "\n"
	}
	return ppm
}
