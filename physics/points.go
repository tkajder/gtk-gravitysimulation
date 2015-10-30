package physics

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

// NewPoint creates a new Point with supplied x and y position
func NewPoint(x float64, y float64) *Point {
	return &Point{X: x, Y: y}
}

// Add returns a new point resulting from following the supplied vector
func (p *Point) Add(v *Vector2D) *Point {
	return NewPoint(p.X+v.X, p.Y+v.Y)
}

// Subtract returns a new point resulting from negatively following the supplied vector
func (p *Point) Subtract(v *Vector2D) *Point {
	return NewPoint(p.X-v.X, p.Y-v.Y)
}

// Distance returns the planar distance between two points
func (p1 *Point) Distance(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2))
}

// DisplacementVector returns the vector from the original point to the supplied point
func (p1 *Point) DisplacementVector(p2 *Point) *Vector2D {
	return NewVector2D(p2.X-p1.X, p2.Y-p1.Y)
}

// String returns the formatted string "Point{X: ..., Y: ...)"
func (p *Point) String() string {
	return fmt.Sprintf("Point{X: %v, Y: %v)", p.X, p.Y)
}
