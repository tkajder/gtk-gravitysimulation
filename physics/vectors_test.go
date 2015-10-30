package physics

import (
	"github.com/tkajder/gravitysimulator/utils"
	"math"
	"reflect"
	"testing"
)

func TestVectorAdd(t *testing.T) {
	t.Parallel()
	cases := []struct {
		v1       *Vector2D
		v2       *Vector2D
		expected *Vector2D
	}{
		{NewVector2D(0, 1), NewVector2D(2, 0), NewVector2D(2, 1)},
		{NewVector2D(5, 3), NewVector2D(2, 4), NewVector2D(7, 7)},
		{NewVector2D(4, 3), NewVector2D(-6, -7), NewVector2D(-2, -4)},
		{NewVector2D(1.25, 1.9), NewVector2D(2.3, 3.01), NewVector2D(3.55, 4.91)},
	}

	for _, c := range cases {
		v3 := c.v1.Add(c.v2)
		if !reflect.DeepEqual(v3, c.expected) {
			t.Errorf("Computing %v + %v = %v - expected %v", c.v1, c.v2, v3, c.expected)
		}
	}
}

func TestVectorSubtract(t *testing.T) {
	t.Parallel()
	cases := []struct {
		v1       *Vector2D
		v2       *Vector2D
		expected *Vector2D
	}{
		{NewVector2D(0, 1), NewVector2D(2, 0), NewVector2D(-2, 1)},
		{NewVector2D(5, 3), NewVector2D(2, 4), NewVector2D(3, -1)},
		{NewVector2D(4, 3), NewVector2D(-6, -7), NewVector2D(10, 10)},
		{NewVector2D(1.25, 2.0), NewVector2D(1.0, 3.5), NewVector2D(0.25, -1.5)},
	}

	for _, c := range cases {
		v3 := c.v1.Subtract(c.v2)
		if !reflect.DeepEqual(v3, c.expected) {
			t.Errorf("Computing %v - %v = %v - expected %v", c.v1, c.v2, v3, c.expected)
		}
	}
}

func TestVectorScalarmul(t *testing.T) {
	t.Parallel()
	cases := []struct {
		v1       *Vector2D
		scalar   float64
		expected *Vector2D
	}{
		{NewVector2D(0, 1), 1.0, NewVector2D(0, 1)},
		{NewVector2D(5, 3), -1.0, NewVector2D(-5, -3)},
		{NewVector2D(4, 3), 0.0, NewVector2D(0, 0)},
		{NewVector2D(0.5, 2.0), 1.5, NewVector2D(0.75, 3.0)},
	}

	for _, c := range cases {
		v2 := c.v1.Scalarmul(c.scalar)
		if !reflect.DeepEqual(v2, c.expected) {
			t.Errorf("Computing %v * %v = %v - expected %v", c.v1, c.scalar, v2, c.expected)
		}
	}
}

func TestVectorDotproduct(t *testing.T) {
	t.Parallel()
	cases := []struct {
		v1       *Vector2D
		v2       *Vector2D
		expected float64
	}{
		{NewVector2D(0, 1), NewVector2D(2, 0), 0},
		{NewVector2D(5, 3), NewVector2D(2, 4), 22},
		{NewVector2D(4, 3), NewVector2D(-6, -7), -45},
		{NewVector2D(1.25, 2.0), NewVector2D(1.0, 3.5), 8.25},
	}

	for _, c := range cases {
		dp := c.v1.Dotproduct(c.v2)
		if dp != c.expected {
			t.Errorf("Computing %v . %v = %v - expected %v", c.v1, c.v2, dp, c.expected)
		}
	}
}

func TestVectorLength(t *testing.T) {
	t.Parallel()
	cases := []struct {
		v        *Vector2D
		expected float64
	}{
		{NewVector2D(0, 1), 1},
		{NewVector2D(4, 3), 5},
		{NewVector2D(8, -6), 10},
		{NewVector2D(0.5, 1.2), 1.3},
	}

	for _, c := range cases {
		l := c.v.Length()
		if l != c.expected {
			t.Errorf("Computing len(%v) = %v - expected %v", c.v, l, c.expected)
		}
	}
}

func TestVectorNormalize(t *testing.T) {
	t.Parallel()
	testprecision := 4
	cases := []struct {
		v        *Vector2D
		expected *Vector2D
	}{
		{NewVector2D(0, 1), NewVector2D(0, 1)},
		{NewVector2D(4, 3), NewVector2D(0.8, 0.6)},
		{NewVector2D(8, -6), NewVector2D(0.8, -0.6)},
		{NewVector2D(0.5, 1.2), NewVector2D(0.3846, 0.9231)},
	}

	for _, c := range cases {
		normv := c.v.Normalize()
		if utils.RoundPrecision(normv.X, testprecision) != c.expected.X || utils.RoundPrecision(normv.Y, testprecision) != c.expected.Y {
			t.Errorf("Computing norm(%v) = %v - expected %v", c.v, normv, c.expected)
		}
	}
}

func TestVectorRotate(t *testing.T) {
	t.Parallel()
	testprecision := 4
	cases := []struct {
		v        *Vector2D
		radians  float64
		expected *Vector2D
	}{
		{NewVector2D(0, 1), 0, NewVector2D(0, 1)},
		{NewVector2D(0, 1), math.Pi / 2, NewVector2D(-1, 0)},
		{NewVector2D(3, 0), math.Pi, NewVector2D(-3, 0)},
		{NewVector2D(2, 2), math.Pi / 4, NewVector2D(0, 2.8284)},
		{NewVector2D(0.5, 1.2), -math.Pi / 3, NewVector2D(1.2892, 0.1670)},
	}

	for _, c := range cases {
		rotv := c.v.Rotate(c.radians)
		if utils.RoundPrecision(rotv.X, testprecision) != c.expected.X || utils.RoundPrecision(rotv.Y, testprecision) != c.expected.Y {
			t.Errorf("Computing rot(%v, %v) = %v - expected %v", c.v, c.radians, rotv, c.expected)
		}
	}
}
func TestVectorInvertX(t *testing.T) {
	t.Parallel()
	cases := []struct {
		v        *Vector2D
		expected *Vector2D
	}{
		{NewVector2D(0, 0), NewVector2D(0, 0)},
		{NewVector2D(1, 0), NewVector2D(-1, 0)},
		{NewVector2D(0, 1), NewVector2D(0, 1)},
		{NewVector2D(-3.214, 5.431), NewVector2D(3.214, 5.431)},
	}

	for _, c := range cases {
		invertedxv := c.v.InvertX()
		if !reflect.DeepEqual(invertedxv, c.expected) {
			t.Errorf("Computing invertx(%v) = %v - expected %v", c.v, invertedxv, c.expected)
		}
	}
}

func TestVectorInvertY(t *testing.T) {
	t.Parallel()
	cases := []struct {
		v        *Vector2D
		expected *Vector2D
	}{
		{NewVector2D(0, 0), NewVector2D(0, 0)},
		{NewVector2D(1, 0), NewVector2D(1, 0)},
		{NewVector2D(0, 1), NewVector2D(0, -1)},
		{NewVector2D(3.214, -5.431), NewVector2D(3.214, 5.431)},
	}

	for _, c := range cases {
		invertedyv := c.v.InvertY()
		if !reflect.DeepEqual(invertedyv, c.expected) {
			t.Errorf("Computing inverty(%v) = %v - expected %v", c.v, invertedyv, c.expected)
		}
	}
}

func TestVectorInvert(t *testing.T) {
	t.Parallel()
	cases := []struct {
		v        *Vector2D
		expected *Vector2D
	}{
		{NewVector2D(0, 0), NewVector2D(0, 0)},
		{NewVector2D(1, 0), NewVector2D(-1, 0)},
		{NewVector2D(0, 1), NewVector2D(0, -1)},
		{NewVector2D(1, -1), NewVector2D(-1, 1)},
		{NewVector2D(-3.214, 5.431), NewVector2D(3.214, -5.431)},
	}

	for _, c := range cases {
		invertedv := c.v.Invert()
		if !reflect.DeepEqual(invertedv, c.expected) {
			t.Errorf("Computing invert(%v) = %v - expected %v", c.v, invertedv, c.expected)
		}
	}
}
