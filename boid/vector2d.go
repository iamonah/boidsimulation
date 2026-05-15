package main

import "math"

type Vector2D struct {
	X float64
	Y float64
}

func (v Vector2D) Add(other Vector2D) Vector2D {
	return Vector2D{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v Vector2D) Subtract(other Vector2D) Vector2D {
	return Vector2D{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v Vector2D) Multiply(scalar float64) Vector2D {
	return Vector2D{
		X: v.X * scalar,
		Y: v.Y * scalar,
	}
}

func (v Vector2D) Divide(scalar float64) Vector2D {
	return Vector2D{
		X: v.X / scalar,
		Y: v.Y / scalar,
	}
}

func (v Vector2D) AddV(d float64) Vector2D {
	return Vector2D{
		X: v.X + d,
		Y: v.Y + d,
	}
}

func (v Vector2D) SubtractV(d float64) Vector2D {
	return Vector2D{
		X: v.X - d,
		Y: v.Y - d,
	}
}
func (v Vector2D) MultiplyV(d float64) Vector2D {
	return Vector2D{
		X: v.X * d,
		Y: v.Y * d,
	}
}
func (v Vector2D) DivideV(d float64) Vector2D {
	return Vector2D{
		X: v.X / d,
		Y: v.Y / d,
	}
}

func (v Vector2D) Limit(lower, max float64) Vector2D {
	return Vector2D{X: math.Min(math.Max(v.X, lower), max), Y: math.Min(math.Max(v.Y, lower), max)}
}

func (v Vector2D) Distance(other Vector2D) float64 {
	//pythagoras theorem for distance between two points in 2D space
	//a^2 + b^2 = c^2	
	return math.Sqrt(math.Pow(v.X-other.X, 2) + math.Pow(v.Y-other.Y, 2))
}
