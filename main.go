package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
	// no need to import player.go golang auto import when running go main.go player.go
)

func main() {
	InitWindow(100, 100, "raylib [core] example - basic window")
	defer CloseWindow()

	SetTargetFPS(60)

	player := Player{
		LoadTexture("assets/player.png"),
		NewRectangle(10, 10, 12, 12),
		NewRectangle(10, 10, 88, 88),
		NewVector2(0, 0),
		0.0,
		White,
		NewVector2(0, 0),
		1.0,
	}

	controls := [4]int32{
		KeyW,
		KeyS,
		KeyA,
		KeyD,
	}

	for !WindowShouldClose() {

		BeginDrawing()

		ClearBackground(RayWhite)

		player = PlayerMove(controls, player)

		player = PlayerUpdate(player)

		PlayerDraw(player)

		EndDrawing()
	}
}
