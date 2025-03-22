package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1980, 1080, "asteroids spin-off game")

	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	//audio initialize, need to use streaming for background music
	rl.InitAudioDevice()
	chillMusic := rl.LoadMusicStream("C:/_dev/Go/cs4953/WeeklyProject7/WeeklyProject7cs4953/670039__seth_makes_sounds__chill-background-music.wav")
	rl.PlayMusicStream(chillMusic)
	rl.SetMusicVolume(chillMusic, 0.2)
	spaceShot := rl.LoadSound("C:/_dev/Go/cs4953/WeeklyProject7/WeeklyProject7cs4953/584198__unfa__weapons-plasma-shot-06.flac")
	rl.SetSoundVolume(spaceShot, 0.6)

	//load the spaceship image
	spaceShipImage := rl.LoadTexture("C:/_dev/Go/cs4953/RaylibAssets/Spaceship.png")
	player := NewSpaceship(spaceShipImage, rl.NewVector2(50, 300), float32(10), float32(100), float32(0))

	//begin drawing
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.SkyBlue)
		//play audio
		rl.UpdateMusicStream(chillMusic)

		//draw the planet and spaceship, probably want to come back and fix the planet
		rl.DrawCircle(int32(rl.GetScreenWidth()/2), int32(rl.GetScreenHeight()/2), 100, rl.Black)
		player.DrawCreature()
		forward, right := float32(0), float32(0)
		//dir := rl.NewVector2(0, 0)
		//basic movement for the spaceship, need to add use of Q and R to look around with spaceship
		if rl.IsKeyDown(rl.KeyW) {
			//dir.Y -= 20
			forward = 20
		}
		if rl.IsKeyDown(rl.KeyS) {
			//dir.Y += 20
			forward = -20
		}
		if rl.IsKeyDown(rl.KeyA) {
			//dir.X -= 20
			right = -20
		}
		if rl.IsKeyDown(rl.KeyD) {
			//dir.X += 20
			right = 20
		}
		//angle the ship
		if rl.IsKeyDown(rl.KeyQ) {
			player.Angle -= 10
		}
		if rl.IsKeyDown(rl.KeyR) {
			player.Angle += 10
		}
		//projectile shot if player presses space
		if rl.IsKeyPressed(rl.KeySpace) {
			rl.PlaySound(spaceShot)
			player.Shoot()
		}

		player.UpdateProjectiles()
		player.DrawProjectiles()
		//player.MoveSpaceship(dir)
		player.MoveSpaceshipWithAngle(forward, right)

		rl.EndDrawing()
	}
}
