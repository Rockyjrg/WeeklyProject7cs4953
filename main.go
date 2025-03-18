package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 400, "Asteroids Spin-off")

	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	//load the spaceship image
	spaceShipImage := rl.LoadTexture("C:/_dev/Go/cs4953/RaylibAssets/Spaceship.png")
	player := NewSpaceship(spaceShipImage, float32(50), float32(300), float32(10), float32(100), float32(0))

	//begin drawing
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.SkyBlue)

		//draw the planet and spaceship, probably want to come back and fix the planet
		rl.DrawCircle(200, 100, 40, rl.Black)
		player.DrawCreature()

		rl.EndDrawing()
	}
}
