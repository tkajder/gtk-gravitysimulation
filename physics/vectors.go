package physics

import (
	"fmt"
	"math"
)

type Vector2D struct {
	X float64
	Y float64
}

// NewVector2D creates a new Vector2D with supplied x and y components
func NewVector2D(x float64, y float64) *Vector2D {
	return &Vector2D{X: x, Y: y}
}

// Add returns a new Vector2D of the addition of two vectors.
func (v1 *Vector2D) Add(v2 *Vector2D) *Vector2D {
	return NewVector2D(v1.X+v2.X, v1.Y+v2.Y)
}

// Subtract returns a new Vector2D of the subtraction of two vectors.
func (v1 *Vector2D) Subtract(v2 *Vector2D) *Vector2D {
	return NewVector2D(v1.X-v2.X, v1.Y-v2.Y)
}

// Scalarmul returns a new Vector2D of the vector components multiplied
// each by the scalar value.
func (v *Vector2D) Scalarmul(scalar float64) *Vector2D {
	return NewVector2D(v.X*scalar, v.Y*scalar)
}

// Dotproduct returns the dot product of two vectors.
func (v1 *Vector2D) Dotproduct(v2 *Vector2D) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

// Length returns the scalar length of the vector.
func (v *Vector2D) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Normalize returns a new Vector2D that is a normalized version
// of the original vector.
func (v *Vector2D) Normalize() *Vector2D {
	normlen := 1.0 / v.Length()
	return NewVector2D(v.X*normlen, v.Y*normlen)
}

// Rotate returns a new Vector2D that is a counterclockwise
// rotation of given radians of the original vector.
func (v *Vector2D) Rotate(radians float64) *Vector2D {
	cr := math.Cos(radians)
	cs := math.Sin(radians)
	return NewVector2D(v.X*cr-v.Y*cs, v.X*cs+v.Y*cr)
}

// InvertX returns a new Vector2D that has an inverted x component
func (v *Vector2D) InvertX() *Vector2D {
	// Do not use the IEEE754 negative zero
	if v.X == 0 {
		return NewVector2D(0, v.Y)
	}
	return NewVector2D(-v.X, v.Y)
}

// InvertY returns a new Vector2D that has an inverted y component
func (v *Vector2D) InvertY() *Vector2D {
	// Do not use the IEEE754 negative zero
	if v.Y == 0 {
		return NewVector2D(v.X, 0)
	}
	return NewVector2D(v.X, -v.Y)
}

// Invert returns a new Vector2D that has both x and y component inverted
func (v *Vector2D) Invert() *Vector2D {
	// Do not use the IEEE754 negative zero
	if v.X == 0 && v.Y == 0 {
		return NewVector2D(0, 0)
	} else if v.X == 0 {
		return NewVector2D(0, -v.Y)
	} else if v.Y == 0 {
		return NewVector2D(-v.X, 0)
	} else {
		return NewVector2D(-v.X, -v.Y)
	}
}

// String returns the formatted string "Vector{X: ..., ...}".
func (v *Vector2D) String() string {
	return fmt.Sprintf("Vector{X: %v, %v}", v.X, v.Y)
}
