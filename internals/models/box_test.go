package models_test

import (
	"menger-sponge-fractal/internals/models"
	"testing"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func TestBox(t *testing.T) {
	box := models.NewBox(0, 0, 0, 1)
	if box.Position != rl.NewVector3(0, 0, 0) {
		t.Error("Box position should be (0, 0, 0)")
	}
	if box.Size != 1 {
		t.Error("Box size should be 1")
	}
}

func TestBoxGenerate(t *testing.T) {
	box := models.NewBox(0, 0, 0, 1)
	boxes := box.Generate()
	if len(boxes) != 20 {
		t.Error("Box should generate 20 boxes")
	}
}
