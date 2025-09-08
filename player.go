package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Sprite    Texture2D
	Dest      Rectangle
	Source    Rectangle
	Origin    Vector2
	Rotation  float32
	Tint      Color
	Direction Vector2
	Speed     float32
}

func PlayerDraw(player Player) {
	DrawTexturePro(
		player.Sprite,
		player.Source,
		player.Dest,
		player.Origin,
		player.Rotation,
		player.Tint,
	)
}

func PlayerMove(controls [4]int32, player Player) Player {
	if IsKeyDown(controls[0]) {
		player.Direction.Y = -1
	} else if IsKeyDown(controls[1]) {
		player.Direction.Y = 1
	} else {
		player.Direction.Y = 0
	}
	if IsKeyDown(controls[2]) {
		player.Direction.X = -1
	} else if IsKeyDown(controls[3]) {
		player.Direction.X = 1
	} else {
		player.Direction.X = 0
	}

	return player
}

func PlayerUpdate(player Player) Player {
	player.Dest.X += player.Direction.X * player.Speed
	player.Dest.Y += player.Direction.Y * player.Speed

	return player
}
