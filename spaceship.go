package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// spaceship(player) struct to draw the spaceship for the game
type Spaceship struct {
	Texture     rl.Texture2D
	Pos         rl.Vector2
	Speed       float32
	Size        float32
	Angle       float32
	Projectiles []Projectile
}

// projectile to shoot out of the spaceship
type Projectile struct {
	Pos   rl.Vector2
	Vel   rl.Vector2
	Speed float32
	Size  float32
	Alive bool
}

// new spaceship using the Spaceship struct
func NewSpaceship(texture rl.Texture2D, newPos rl.Vector2, speed, size, angle float32) Spaceship {
	return Spaceship{
		Texture: texture,
		Pos:     newPos,
		Speed:   speed,
		Size:    size,
		Angle:   angle,
	}
}

// draw the spaceship to the screen
func (s Spaceship) DrawCreature() {
	scale := s.Size / float32(s.Texture.Width)
	//rl.DrawTextureEx(s.Texture, s.Pos, s.Angle, scale, rl.White)
	// Define the source rectangle (entire texture)
	srcRect := rl.Rectangle{X: 0, Y: 0, Width: float32(s.Texture.Width), Height: float32(s.Texture.Height)}

	// Define the destination rectangle (scaled position & size)
	dstRect := rl.Rectangle{
		X:      s.Pos.X,
		Y:      s.Pos.Y,
		Width:  float32(s.Texture.Width) * scale,
		Height: float32(s.Texture.Height) * scale,
	}

	// Set origin to center of the texture
	origin := rl.Vector2{X: dstRect.Width / 2, Y: dstRect.Height / 2}

	// Draw with correct rotation and origin
	rl.DrawTexturePro(s.Texture, srcRect, dstRect, origin, s.Angle, rl.White)
}

func (s *Spaceship) MoveSpaceshipWithAngle(forward, right float32) {
	rad := float64(s.Angle) * (math.Pi / 180) // Convert to radians

	// Forward movement direction
	forwardVec := rl.NewVector2(float32(math.Sin(rad)), -float32(math.Cos(rad)))

	// Right movement direction (perpendicular to forward)
	radRight := float64(s.Angle+90) * (math.Pi / 180)
	rightVec := rl.NewVector2(float32(math.Sin(radRight)), float32(math.Cos(radRight)))

	// Calculate movement using forward and right inputs
	movement := rl.Vector2Add(rl.Vector2Scale(forwardVec, forward*s.Speed*rl.GetFrameTime()), rl.Vector2Scale(rightVec, right*s.Speed*rl.GetFrameTime()))

	// Apply movement to spaceship
	s.Pos = rl.Vector2Add(s.Pos, movement)
}

// create a new projectile from the spaceship's position and angle
func NewProjectile(pos rl.Vector2, angle float32, speed float32) Projectile {
	rad := float64(angle) * (math.Pi / 180)

	vel := rl.NewVector2(float32(math.Sin(rad)), -float32(math.Cos(rad)))

	return Projectile{
		Pos:   pos,
		Vel:   rl.Vector2Scale(vel, speed),
		Speed: speed,
		Size:  5,
		Alive: true,
	}
}

// move projectile in direction of its velocity
func (p *Projectile) Update() {
	p.Pos = rl.Vector2Add(p.Pos, rl.Vector2Scale(p.Vel, rl.GetFrameTime()))
}

// draw projectile
func (p Projectile) Draw() {
	if p.Alive {
		rl.DrawCircle(int32(p.Pos.X), int32(p.Pos.Y), p.Size, rl.Red)
	}
}

// shoot projectiles
func (s *Spaceship) Shoot() {
	projectileStart := rl.Vector2Add(s.Pos, rl.NewVector2(0, -s.Size/2))

	s.Projectiles = append(s.Projectiles, NewProjectile(projectileStart, s.Angle, 500))
}

func (s *Spaceship) UpdateProjectiles() {
	for i := range s.Projectiles {
		s.Projectiles[i].Update()
	}
}

func (s *Spaceship) DrawProjectiles() {
	for _, p := range s.Projectiles {
		p.Draw()
	}
}
