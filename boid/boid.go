package main

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	Position Vector2D
	Velocity Vector2D
	id       int
}

func (b *Boid) calcAcceleration() Vector2D {
	//find neighbors and calculate acceleration based on separation, alignment, and cohesion
	uper, lower := b.Position.AddV(viewRadius), b.Position.AddV(-viewRadius)
	avgPosition, avgVelocity := Vector2D{X: 0, Y: 0}, Vector2D{X: 0, Y: 0}
	neighborCount := 0.0

	lock.Lock()
	for i := math.Max(lower.X, 0); i <= math.Min(uper.X, screenWidth); i++ {
		for j := math.Max(lower.Y, 0); j <= math.Min(uper.Y, screenHeight); j++ {
			if otherBoidId := boidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.id {
				//check if the other boid is within the view radius
				if dist := boids[otherBoidId].Position.Distance(b.Position); dist < viewRadius {
					neighborCount++
					avgVelocity = avgVelocity.Add(boids[otherBoidId].Velocity)
					avgPosition = avgPosition.Add(boids[otherBoidId].Position)
				}
			}
		}
	}
	lock.Unlock()
	accel := Vector2D{X: 0, Y: 0}
	if neighborCount > 0 {
		avgPosition, avgVelocity = avgPosition.DivideV(neighborCount), avgVelocity.DivideV(neighborCount)
		accelAlignment := avgVelocity.Subtract(b.Velocity).MultiplyV(adjustmentFactor)
		accCohesion := avgPosition.Subtract(b.Position).MultiplyV(adjustmentFactor)
		accel = accel.Add(accelAlignment).Add(accCohesion)
	}
	return accel
}

func (b *Boid) moveOne() {
	acceleration := b.calcAcceleration()
	lock.Lock()
	b.Velocity = b.Velocity.Add(acceleration)
	//remove old position or location from map
	boidMap[int(b.Position.X)][int(b.Position.Y)] = -1
	b.Position = b.Position.Add(b.Velocity)
	//update new location
	boidMap[int(b.Position.X)][int(b.Position.Y)] = b.id
	next := b.Position.Add(b.Velocity)
	if next.X < 0 || next.X > screenWidth {
		b.Velocity = Vector2D{-b.Velocity.X, b.Velocity.Y}
	}
	if next.Y < 0 || next.Y > screenHeight {
		b.Velocity = Vector2D{b.Velocity.X, -b.Velocity.Y}
	}
	lock.Unlock()
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func createBoid(id int) {
	boid := Boid{
		Position: Vector2D{X: rand.Float64() * screenWidth, Y: rand.Float64() * screenHeight},
		Velocity: Vector2D{X: rand.Float64()*2 - 1, Y: rand.Float64()*2 - 1},
		id:       id,
	}
	boids[id] = &boid
	boidMap[int(boid.Position.X)][int(boid.Position.Y)] = boid.id
	go boid.start()
}
