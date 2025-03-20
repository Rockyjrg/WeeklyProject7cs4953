package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 400, "Asteroids Spin-off")

	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	//audio initialize, need to use streaming for background music
	rl.InitAudioDevice()
	chillMusic := rl.LoadMusicStream("C:/_dev/Go/cs4953/WeeklyProject7/WeeklyProject7cs4953/670039__seth_makes_sounds__chill-background-music.wav")
	rl.PlayMusicStream(chillMusic)
	rl.SetMusicVolume(chillMusic, 0.2)

	//load the spaceship image
	spaceShipImage := rl.LoadTexture("C:/_dev/Go/cs4953/RaylibAssets/Spaceship.png")
	player := NewSpaceship(spaceShipImage, float32(50), float32(300), float32(10), float32(100), float32(0))

	//begin drawing
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.SkyBlue)
		//play audio
		rl.UpdateMusicStream(chillMusic)

		//draw the planet and spaceship, probably want to come back and fix the planet
		rl.DrawCircle(200, 100, 40, rl.Black)
		player.DrawCreature()

		//basic movement for the spaceship, need to add use of Q and R to look around with spaceship
		if rl.IsKeyDown(rl.KeyW) {
			player.MoveSpaceship(0, -20)
		}
		if rl.IsKeyDown(rl.KeyA) {
			player.MoveSpaceship(-20, 0)
		}
		if rl.IsKeyDown(rl.KeyS) {
			player.MoveSpaceship(0, 20)
		}
		if rl.IsKeyDown(rl.KeyD) {
			player.MoveSpaceship(20, 0)
		}

		rl.EndDrawing()
	}
}
