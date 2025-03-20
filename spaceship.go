package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// spaceship(player) struct to draw the spaceship for the game
type Spaceship struct {
	Texture rl.Texture2D
	Xpos    float32
	Ypos    float32
	Speed   float32
	Size    float32
	Angle   float32
}

// new spaceship using the Spaceship struct
func NewSpaceship(texture rl.Texture2D, xpos, ypos, speed, size, angle float32) Spaceship {
	return Spaceship{
		Texture: texture,
		Xpos:    xpos,
		Ypos:    ypos,
		Speed:   speed,
		Size:    size,
		Angle:   angle,
	}
}

// draw the spaceship to the screen
func (s Spaceship) DrawCreature() {
	scale := s.Size / float32(s.Texture.Width)
	rl.DrawTextureEx(s.Texture, rl.Vector2{X: s.Xpos, Y: s.Ypos}, s.Angle, scale, rl.White)
}

// move spaceship function
func (s *Spaceship) MoveSpaceship(xOffset, yOffset float32) {
	s.Xpos += xOffset
	s.Ypos += yOffset

	//top and bottom boundaries
	if s.Ypos < 0 {
		s.Ypos = 0
	}
	if s.Ypos+s.Size > float32(rl.GetScreenHeight()) {
		s.Ypos = float32(rl.GetScreenHeight()) - s.Size
	}
	//left and right boundaries
	if s.Xpos < 0 {
		s.Xpos = 0
	}
	if s.Xpos+s.Size > float32(rl.GetScreenWidth()) {
		s.Xpos = float32(rl.GetScreenWidth()) - s.Size
	}
}
