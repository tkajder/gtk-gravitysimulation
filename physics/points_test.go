package physics

import (
	"reflect"
	"testing"
)

func TestPointAdd(t *testing.T) {
	t.Parallel()
	cases := []struct {
		point    *Point
		vector   *Vector2D
		expected *Point
	}{
		{NewPoint(0, 0), NewVector2D(1, 1), NewPoint(1, 1)},
		{NewPoint(1, 1), NewVector2D(2, 3), NewPoint(3, 4)},
		{NewPoint(1, 2), NewVector2D(-4, -7), NewPoint(-3, -5)},
		{NewPoint(1.25, 2.5), NewVector2D(-0.75, 3.75), NewPoint(0.5, 6.25)},
	}

	for _, c := range cases {
		p := c.point.Add(c.vector)
		if !reflect.DeepEqual(p, c.expected) {
			t.Errorf("Computing %v + %v = %v - expected %v", c.point, c.vector, p, c.expected)
		}
	}
}

func TestPointSubtract(t *testing.T) {
	t.Parallel()
	cases := []struct {
		point    *Point
		vector   *Vector2D
		expected *Point
	}{
		{NewPoint(0, 0), NewVector2D(1, 1), NewPoint(-1, -1)},
		{NewPoint(8, 10), NewVector2D(3, 8), NewPoint(5, 2)},
		{NewPoint(1, 2), NewVector2D(-4, -7), NewPoint(5, 9)},
		{NewPoint(1.25, 2.5), NewVector2D(-0.75, 3.5), NewPoint(2, -1)},
	}

	for _, c := range cases {
		p := c.point.Subtract(c.vector)
		if !reflect.DeepEqual(p, c.expected) {
			t.Errorf("Computing %v - %v = %v - expected %v", c.point, c.vector, p, c.expected)
		}
	}
}

func TestPointDistance(t *testing.T) {
	t.Parallel()
	cases := []struct {
		p1       *Point
		p2       *Point
		expected float64
	}{
		{NewPoint(0, 0), NewPoint(0, 0), 0},
		{NewPoint(0, 0), NewPoint(0, 1), 1},
		{NewPoint(0, 0), NewPoint(3, 4), 5},
		{NewPoint(0, 0), NewPoint(-3, -4), 5},
		{NewPoint(3, 4), NewPoint(6, 8), 5},
	}

	for _, c := range cases {
		distance := c.p1.Distance(c.p2)
		if distance != c.expected {
			t.Errorf("Computing distance between %v and %v = %v - expected %v", c.p1, c.p2, distance, c.expected)
		}
	}
}

func TestDisplacementVector(t *testing.T) {
	t.Parallel()
	cases := []struct {
		p1       *Point
		p2       *Point
		expected *Vector2D
	}{
		{NewPoint(0, 0), NewPoint(0, 0), NewVector2D(0, 0)},
		{NewPoint(0, 0), NewPoint(0, 1), NewVector2D(0, 1)},
		{NewPoint(0, 0), NewPoint(3, 4), NewVector2D(3, 4)},
		{NewPoint(0, 0), NewPoint(-3, -4), NewVector2D(-3, -4)},
		{NewPoint(3, 4), NewPoint(6, 8), NewVector2D(3, 4)},
	}

	for _, c := range cases {
		dv := c.p1.DisplacementVector(c.p2)
		if !reflect.DeepEqual(dv, c.expected) {
			t.Errorf("Computing displacement vector between %v and %v = %v - expected %v", c.p1, c.p2, dv, c.expected)
		}
	}
}
