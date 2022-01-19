package lights

import (
	"bolijollo/rayos/color"
	mat "bolijollo/rayos/materials"
	p "bolijollo/rayos/primitives"
	"fmt"
	"testing"
)

func TestNightningCalculation(t *testing.T) {
	eyev := p.Vector(0, 0, -1)
	point := p.Point(0, 0, 0)
	normalv := p.Vector(0, 0, -1)
	light := CreatePointLight(p.Point(0, 10, -10), color.White)
	m := mat.DefaultMaterial()
	result := Lighting(m, &light, point, normalv, eyev)
	fmt.Println("result es ahora veamos:", result, "y esperamos:", 0.7364*255)
}
