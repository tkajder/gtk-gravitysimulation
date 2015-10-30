package physics

import (
	"fmt"
	"math"
)

type Entity struct {
	Mass         float64
	Position     *Point
	Velocity     *Vector2D
	Acceleration *Vector2D
}

// Made up gravity so reactions on the seconds scale at small range are fun to watch
const G float64 = 6.67834E2

// NewEntity returns a new Entity struct from the provided float values
func NewEntity(mass float64, posx float64, posy float64, velx float64, vely float64, accelx float64, accely float64) *Entity {
	return &Entity{Mass: mass, Position: NewPoint(posx, posy), Velocity: NewVector2D(velx, vely), Acceleration: NewVector2D(accelx, accely)}
}

// GravitationalForce returns the gravitational force between two entites based on both entities masses and distance
func (e1 *Entity) GravitationalForce(e2 *Entity) float64 {
	return (G * e1.Mass * e2.Mass) / math.Pow(e1.Distance(e2), 2)
}

// Update updates the position and velocity of the Entity for a given time tick
func (e1 *Entity) Update(time float64) {
	e1.Position = e1.Position.Add(e1.Velocity.Scalarmul(time))
	e1.Velocity = e1.Velocity.Add(e1.Acceleration.Scalarmul(time))
}

// UpdateGravity updates the acceleration of the Entity based on the aggregate gravitational acceleration of the given entities slice upon the entitiy
func (e1 *Entity) UpdateGravitationalAcceleration(entities []*Entity) {

	// Reset acceleration to 0-vector
	e1.Acceleration = NewVector2D(0, 0)

	for _, e2 := range entities {

		// Entities should exert no gravity upon themselves - skip
		if e1 == e2 {
			continue
		}

		// Using gravitational force find the gravitational acceleration vector from e1 to e1
		gforce := e1.GravitationalForce(e2)
		gaccelscalar := gforce / e1.Mass
		normal := e1.Position.DisplacementVector(e2.Position).Normalize()
		e1.Acceleration = e1.Acceleration.Add(normal.Scalarmul(gaccelscalar))
	}
}

// Distance returns the positional distance between two entities
func (e1 *Entity) Distance(e2 *Entity) float64 {
	return e1.Position.Distance(e2.Position)

}

// String returns the formatted string "Entity{Mass: ..., Position: ..., Velocity: ..., Acceleration: ...}"
func (e *Entity) String() string {
	return fmt.Sprintf("Entity{Mass: %v, Position: %v, Velocity: %v, Acceleration: %v}", e.Mass, e.Position, e.Velocity, e.Acceleration)
}
