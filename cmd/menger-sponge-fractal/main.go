package main

import (
	"math"
	"menger-sponge-fractal/internals/models"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 400
	screenHeight = 400
)

var (
	// backgroundColor = rl.NewColor(51, 51, 51, 255)
	backgroundColor = rl.Green
)

func main() {
	rl.InitWindow(400, 400, "menger-sponge-fractal")
	rl.SetTargetFPS(60)
	// center the camera
	camera := rl.Camera3D{
		Position:   rl.NewVector3(5, 5, 5),
		Target:     rl.NewVector3(0, 0, 0),
		Up:         rl.NewVector3(0, 3, 0),
		Fovy:       45,
		Projection: rl.CameraPerspective,
	}
	sponge := []models.Box{}
	box := models.NewBox(0, 0, 0, 2)
	sponge = append(sponge, box)
	var rotationAngle float32
	for !rl.WindowShouldClose() {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			var newSponge []models.Box
			for _, b := range sponge {
				newSponge = append(newSponge, b.Generate()...)
			}
			sponge = newSponge
		}
		rotationAngle += rl.GetFrameTime() * 90
		camera.Position = rl.NewVector3(
			5*float32(math.Cos(float64(rotationAngle)*rl.Deg2rad)),
			5*float32(math.Sin(float64(rotationAngle)*rl.Deg2rad)),
			camera.Position.Z,
		)
		rl.UpdateCamera(&camera, rl.CameraOrbital)
		rl.BeginDrawing()
		rl.ClearBackground(backgroundColor)
		rl.BeginMode3D(camera)
		for _, b := range sponge {
			b.Draw()
		}
		rl.EndMode3D()
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

