package main

import (
	"fmt"
	"os"
	"rayos/color"

	"rayos/canvas"
	p "rayos/math"
)

type particle struct {
	Position p.Tuple
	Velocity p.Tuple
}

func main() {
	ball := particle{
		Position: p.Vector(0, 0, 0),
		Velocity: p.Vector(0.1, 0.5, 0)}

	friction_coefficient := 0.01
	gravity := p.Vector(0, -0.1, 0)
	wind := p.Vector(-0.005, 0.01, 0)
	file := "test.ppm"

	// Light_Blue := Verde.Add(&Azul)

	Trapo := canvas.NewCanvas(800, 600)

	for ball.Position.Y >= 0 {
		fmt.Println(ball.Position)
		Trapo.Pixels[int(ball.Position.X)][int(ball.Position.Y)] = color.NewColor(0, 0, 255)
		ball.Velocity = p.Times(&ball.Velocity, 1-friction_coefficient)
		ball.Velocity = p.Add(&ball.Velocity, &gravity)
		ball.Velocity = p.Add(&ball.Velocity, &wind)
		ball.Position = p.Add(&ball.Position, &ball.Velocity)

	}

	// Verde := color.NewColor(0, 255, 0)
	// Azul := color.NewColor(0, 0, 255)

	data := Trapo.WritePPM()
	err := os.WriteFile(file, []byte(data), 0644)
	if err != nil {
		fmt.Println(err)
	}

}
