package main

import (
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Asteroid struct {
	Pos   rl.Vector2
	Vel   rl.Vector2
	Speed float32
	Size  float32
	Color rl.Color
}

func NewAsteroid(planetPos rl.Vector2, minDist float32) Asteroid {
	var pos rl.Vector2

	//ensure asteroid spawns at a distance away from the planet
	for {
		pos = rl.NewVector2(rand.Float32()*float32(rl.GetScreenWidth()), rand.Float32()*float32(rl.GetScreenHeight()))
		dist := rl.Vector2Distance(pos, planetPos)
		if dist > minDist {
			break
		}
	}

	//velocity towards planet
	dir := rl.Vector2Subtract(planetPos, pos)
	dir = rl.Vector2Normalize(dir)
	speed := rand.Float32() * 30

	asteroidColor := rl.NewColor(uint8(rand.IntN(256)), uint8(rand.IntN(256)), uint8(rand.IntN(256)), 255)

	return Asteroid{
		Pos:   pos,
		Vel:   rl.Vector2Scale(dir, speed),
		Speed: speed,
		Size:  30,
		Color: asteroidColor,
	}
}
func (a *Asteroid) Update() {
	a.Pos = rl.Vector2Add(a.Pos, rl.Vector2Scale(a.Vel, rl.GetFrameTime()))
}

func (a Asteroid) Draw() {
	rl.DrawCircle(int32(a.Pos.X), int32(a.Pos.Y), a.Size, a.Color)
}

func NewSmallAsteroid(pos rl.Vector2, speed float32) Asteroid {
	//random direction
	dir := rl.NewVector2(rand.Float32()*2-1, rand.Float32()*2-1)
	dir = rl.Vector2Normalize(dir)

	//assign smaller value and color if it's a small asteroid
	//asteroidColor := rl.NewColor(rl.Yellow)

	return Asteroid{
		Pos:   pos,
		Vel:   rl.Vector2Scale(dir, speed),
		Speed: speed,
		Size:  15,
		Color: rl.Yellow,
	}
}
