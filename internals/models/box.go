package models

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Box struct {
	Position rl.Vector3
	Radius   float32
}

func NewBox(x, y, z, size float32) Box {
	return Box{
		Position: rl.NewVector3(x, y, z),
		Radius:   size,
	}
}

func (b *Box) Generate() []Box {
	var boxes []Box
	var x, y, z float32
	for x = -1; x < 2; x++ {
		for y = -1; y < 2; y++ {
			for z = -1; z < 2; z++ {
				sum := math.Abs(float64(x)) + math.Abs(float64(y)) + math.Abs(float64(z))
				newRadius := b.Radius / 3
				if sum > 1 {
					pos := b.Position
					b := NewBox(pos.X+x*newRadius, pos.Y+y*newRadius, pos.Z+z*newRadius, newRadius)
					boxes = append(boxes, b)
				}
			}
		}
	}
	return boxes
}

func (b *Box) Draw() {
	rl.PushMatrix()
	rl.DrawCube(b.Position, b.Radius, b.Radius, b.Radius, rl.White)
	rl.DrawCubeWires(b.Position, b.Radius, b.Radius, b.Radius, rl.Black)
	rl.PopMatrix()
}
