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
	upper, lower := b.Position.AddV(viewRadius), b.Position.AddV(-viewRadius)
	avgPosition, avgVelocity, separation := Vector2D{X: 0, Y: 0}, Vector2D{X: 0, Y: 0}, Vector2D{X: 0, Y: 0}
	neighborCount := 0.0

	rLock.RLock()
	for i := math.Max(lower.X, 0); i <= math.Min(upper.X, screenWidth); i++ {
		for j := math.Max(lower.Y, 0); j <= math.Min(upper.Y, screenHeight); j++ {
			if otherBoidId := boidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != boidId(b.id) {
				//check if the other boid is within the view radius
				if dist := boids[otherBoidId].Position.Distance(b.Position); dist < viewRadius {
					neighborCount++
					avgVelocity = avgVelocity.Add(boids[otherBoidId].Velocity)
					avgPosition = avgPosition.Add(boids[otherBoidId].Position)
					separation = separation.Add(b.Position.Subtract(boids[otherBoidId].Position).DivideV(dist))
				}
			}
		}
	}
	rLock.RUnlock()

	accel := Vector2D{X: b.borderBounce(b.Position.X, screenWidth), Y: b.borderBounce(b.Position.Y, screenHeight)}
	if neighborCount > 0 {
		avgPosition, avgVelocity = avgPosition.DivideV(neighborCount), avgVelocity.DivideV(neighborCount)
		accelAlignment := avgVelocity.Subtract(b.Velocity).MultiplyV(adjustmentFactor)
		accCohesion := avgPosition.Subtract(b.Position).MultiplyV(adjustmentFactor)
		accSeperation := separation.MultiplyV(adjustmentFactor)
		accel = accel.Add(accelAlignment).Add(accCohesion).Add(accSeperation)
	}
	return accel
}

func (b *Boid) borderBounce(position, maxborder float64) float64 {
	if position < viewRadius {
		return 1 / position
	} else if position > maxborder-viewRadius {
		return 1 / (position - maxborder)
	}
	return 0
}

func (b *Boid) moveOne() {
	acceleration := b.calcAcceleration()
	rLock.Lock()
	b.Velocity = b.Velocity.Add(acceleration).Limit(-1, 1)
	//remove old position or location from map
	boidMap[int(b.Position.X)][int(b.Position.Y)] = -1
	b.Position = b.Position.Add(b.Velocity)
	//update new location
	boidMap[int(b.Position.X)][int(b.Position.Y)] = boidId(b.id)
	rLock.Unlock()
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
	boidMap[int(boid.Position.X)][int(boid.Position.Y)] = boidId(boid.id)
	go boid.start()
}
