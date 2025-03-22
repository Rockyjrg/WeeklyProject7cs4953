package main

import (
	"fmt"

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
	spaceShot := rl.LoadSound("C:/_dev/Go/cs4953/WeeklyProject7/WeeklyProject7cs4953/509070__tripjazz__old-school-pew-pew-3.wav")
	rl.SetSoundVolume(spaceShot, 0.8)

	asteroidDestroySound := rl.LoadSound("C:/_dev/Go/cs4953/WeeklyProject7/WeeklyProject7cs4953/630590__vinni_r__bone-break_4.wav")
	rl.SetSoundVolume(asteroidDestroySound, 1.0)

	//load the spaceship image
	spaceShipImage := rl.LoadTexture("C:/_dev/Go/cs4953/RaylibAssets/Spaceship.png")
	player := NewSpaceship(spaceShipImage, rl.NewVector2(float32(rl.GetScreenWidth()/2), 900), float32(10), float32(100), float32(0))

	var asteroids []Asteroid
	var planetPos = rl.NewVector2(float32(rl.GetScreenWidth()/2), float32(rl.GetScreenHeight()/2))

	//timer for spawning of asteroids
	spawnTimer := float32(0)
	spawnRate := float32(2)

	var planetHealth int = 20
	var planetRadius float32 = 100

	//begin drawing
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.SkyBlue)
		//play audio
		rl.UpdateMusicStream(chillMusic)

		//check if game is over
		if planetHealth <= 0 {
			rl.DrawText("GAME OVER", int32(rl.GetScreenWidth())/2-100, int32(rl.GetScreenHeight())/2, 50, rl.Red)
			rl.DrawText("Press R to Restart", int32(rl.GetScreenWidth())/2-130, int32(rl.GetScreenHeight())/2+150, 20, rl.White)

			if rl.IsKeyPressed(rl.KeyR) {
				planetHealth = 20
				asteroids = []Asteroid{}
				player.Pos = rl.NewVector2(float32(rl.GetScreenWidth()/2), 900)
			}

			rl.EndDrawing()
			continue
		}

		//draw the planet and spaceship, probably want to come back and fix the planet
		rl.DrawCircle(int32(planetPos.X), int32(planetPos.Y), 100, rl.Black)
		player.DrawCreature()

		//spawn asteroids at intervals
		spawnTimer += rl.GetFrameTime()
		if spawnTimer >= spawnRate {
			asteroids = append(asteroids, NewAsteroid(planetPos, 300))
			spawnTimer = 0
		}

		//check collision and update
		newAsteroids := []Asteroid{}
		for _, asteroid := range asteroids {
			asteroid.Update()

			if rl.Vector2Distance(asteroid.Pos, planetPos) <= planetRadius+asteroid.Size {
				planetHealth-- //reduce planet health
				continue       //skip to next asteroid
			}

			asteroid.Draw()
			newAsteroids = append(newAsteroids, asteroid)
		}
		asteroids = newAsteroids //update asteroids

		//display planet health
		rl.DrawText(fmt.Sprintf("Planet Health: %d", planetHealth), 10, 10, 20, rl.Red)

		forward, right := float32(0), float32(0)
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
