package lights

import (
	"bolijollo/rayos/color"
	mat "bolijollo/rayos/materials"
	p "bolijollo/rayos/primitives"
	"fmt"
	"testing"
)

func TestLightningCalculation(t *testing.T) {
	eyev := p.Vector(0, 0, -1)
	point := p.Point(0, 0, 0)
	normalv := p.Vector(0, 0, -1)
	light := CreatePointLight(p.Point(0, 10, -10), color.White)
	m := mat.DefaultMaterial()
	result := Lighting(m, &light, point, normalv, eyev, false)
	fmt.Println("result es ahora veamos:", result, "y esperamos:", 0.7364*255)
}

func TestLightningWithSurfaceinShadows(t *testing.T) {
	eyev := p.Vector(0, 0, -1)
	point := p.Point(0, 0, 0)
	normalv := p.Vector(0, 0, -1)
	in_shadow := true
	light := CreatePointLight(p.Point(0, 0, -10), color.White)
	got := Lighting(mat.DefaultMaterial(), &light, point, normalv, eyev, in_shadow)
	want := color.NewColor(0.1, 0.1, 0.1).Scale(255)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
