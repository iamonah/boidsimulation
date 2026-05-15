package main

import (
	"image/color"
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

type boidId int

const (
	screenWidth      = 640.0
	screenHeight     = 360.0
	boidCount        = 500
	viewRadius       = 13
	adjustmentFactor = 0.015
)

var (
	green   = color.RGBA{10, 255, 50, 255}
	boids   [boidCount]*Boid
	boidMap [screenWidth + 1][screenHeight + 1]boidId
	rLock    sync.RWMutex
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, boid := range boids {
		screen.Set(int(boid.Position.X+1), int(boid.Position.Y), green)
		screen.Set(int(boid.Position.X-1), int(boid.Position.Y), green)
		screen.Set(int(boid.Position.X), int(boid.Position.Y-1), green)
		screen.Set(int(boid.Position.X), int(boid.Position.Y+1), green)

	}
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHeight
}

func main() {
	for i, row := range boidMap {
		for j := range row {
			boidMap[i][j] = -1
		}
	}
	for i := 0; i < boidCount; i++ {
		createBoid(i)
	}

	ebiten.SetWindowSize(int(screenWidth*2), int(screenHeight*2))
	ebiten.SetWindowTitle("Boids Simulation")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
