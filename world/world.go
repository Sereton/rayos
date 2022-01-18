package world

import (
	"bolijollo/rayos/lights"
	"bolijollo/rayos/shapes"
)

type World struct {
	Objects []*shapes.Shape      // List of objects in the world
	Lights  []*lights.PointLight // List of lights in the world
}

func CreateWorld() *World {
	return &World{
		Objects: make([]*shapes.Shape, 0),
		Lights:  make([]*lights.PointLight, 0),
	}
}

func (w *World) AddObject(obj shapes.Shape) {
	w.Objects = append(w.Objects, &obj)
}

func (w *World) AddLight(light *lights.PointLight) {
	w.Lights = append(w.Lights, light)
}
