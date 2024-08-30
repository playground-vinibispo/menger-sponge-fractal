package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 400
	screenHeight = 400
)

var (
	backgroundColor = rl.NewColor(51, 51, 51, 255)
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
	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraOrbital)
		rl.BeginDrawing()
		rl.ClearBackground(backgroundColor)
		rl.BeginMode3D(camera)
		rl.DrawCube(rl.NewVector3(0, 0, 0), 2, 2, 2, backgroundColor)
		rl.DrawCubeWires(rl.NewVector3(0, 0, 0), 2, 2, 2, rl.RayWhite)
		rl.EndMode3D()
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

