package physics

import (
	"github.com/tkajder/gravitysimulator/utils"
	"reflect"
	"testing"
)

func TestEntityGravitationalForce(t *testing.T) {
	t.Parallel()
	testprecision := 3
	cases := []struct {
		e1       *Entity
		e2       *Entity
		expected float64
	}{
		{NewEntity(1, 0, 0, 0, 0, 0, 0), NewEntity(1, 1, 0, 0, 0, 0, 0), 667.834},
		{NewEntity(2, 0, 0, 0, 0, 0, 0), NewEntity(2, 2, 0, 0, 0, 0, 0), 667.834},
		{NewEntity(3, 0, 0, 0, 0, 0, 0), NewEntity(6, 5, 0, 0, 0, 0, 0), 480.840},
		{NewEntity(4.2, 0, 0, 0, 0, 0, 0), NewEntity(13.13, 6.841, 0, 0, 0, 0, 0), 786.943},
	}

	for _, c := range cases {
		force := c.e1.GravitationalForce(c.e2)
		if utils.RoundPrecision(force, testprecision) != c.expected {
			t.Errorf("Computing %v.GravitationalForce(%v) = %v - expected %v", c.e1, c.e2, force, c.expected)
		}
	}
}

func TestEntityUpdate(t *testing.T) {
	t.Parallel()
	cases := []struct {
		entity   *Entity
		time     float64
		expected *Entity
	}{
		{NewEntity(1, 0, 0, 0, 0, 0, 0), 1.0, NewEntity(1, 0, 0, 0, 0, 0, 0)},
		{NewEntity(10, 0, 0, 1, 0, 0, 0), 1.0, NewEntity(10, 1, 0, 1, 0, 0, 0)},
		{NewEntity(100, 0, 0, 1, 0, 1, 0), 1.0, NewEntity(100, 1, 0, 2, 0, 1, 0)},
		{NewEntity(2, 0, 0, 1, 1, 1, 1), 0.1, NewEntity(2, 0.1, 0.1, 1.1, 1.1, 1, 1)},
		{NewEntity(5.2, 1.3, -9.1, 14.1, -23, -1, 1), 1.5, NewEntity(5.2, 22.45, -43.6, 12.6, -21.5, -1, 1)},
	}

	for _, c := range cases {
		c.entity.Update(c.time)
		if !reflect.DeepEqual(c.entity, c.expected) {
			t.Errorf("Computting update of %v seconds got %v - expected %v", c.time, c.entity, c.expected)
		}
	}
}

func TestEntityUpdateGravitationalAcceleration(t *testing.T) {
	t.Parallel()
	testprecision := 4
	cases := []struct {
		entity   *Entity
		entities []*Entity
		expected *Entity
	}{
		{NewEntity(1, 0, 0, 0, 0, 0, 0), []*Entity{NewEntity(1E12, 1, 0, 0, 0, 0, 0)}, NewEntity(1, 0, 0, 0, 0, 6.67834E14, 0)},
		{NewEntity(1e20, 0, 0, 0, 0, 0, 0), []*Entity{NewEntity(1e20, 1, 0, 0, 0, 0, 0), NewEntity(1e20, -1, 0, 0, 0, 0, 0)}, NewEntity(1e20, 0, 0, 0, 0, 0, 0)},
		{NewEntity(1, 0, 0, 0, 0, 0, 0), []*Entity{NewEntity(1e12, 1, 0, 0, 0, 0, 0), NewEntity(2e12, 0, 1, 0, 0, 0, 0), NewEntity(1e12, -1, 0, 0, 0, 0, 0)}, NewEntity(1, 0, 0, 0, 0, 0, 2*6.67834E14)},
	}

	for _, c := range cases {
		c.entity.UpdateGravitationalAcceleration(c.entities)
		if accelerationsinequal(c.entity, c.expected, testprecision) {
			t.Errorf("Computing update: got %v - expected %v", c.entity, c.expected)
		}
	}
}

func accelerationsinequal(e1 *Entity, e2 *Entity, testprecision int) bool {
	xinequal := utils.RoundPrecision(e1.Acceleration.X, testprecision) != e2.Acceleration.X
	yinequal := utils.RoundPrecision(e1.Acceleration.Y, testprecision) != e2.Acceleration.Y
	return xinequal || yinequal
}

func TestEntityDistance(t *testing.T) {
	t.Parallel()
	testprecision := 4
	cases := []struct {
		e1       *Entity
		e2       *Entity
		expected float64
	}{
		{NewEntity(1, 0, 0, 0, 0, 0, 0), NewEntity(1, 1, 0, 0, 0, 0, 0), 1},
		{NewEntity(10, 0, 0, 0, 0, 0, 0), NewEntity(1, 3, 4, 0, 0, 0, 0), 5},
		{NewEntity(25, 15, 16, 19, 12, 31, 31), NewEntity(12, -10, -44, -1, -44, 33, -1), 65},
		{NewEntity(1.3, 2.55, 1.22, 5.32, 233, 3.154, 43.1245), NewEntity(1.5467, 2.356, 2.15123, 0.33, 0.23, 1.2, 2.5), 0.9512},
	}

	for _, c := range cases {
		distance := utils.RoundPrecision(c.e1.Distance(c.e2), testprecision)
		if distance != c.expected {
			t.Errorf("Computing distance from %v to %v: got %v - expected %v", c.e1, c.e2, distance, c.expected)
		}
	}
}
