package utils

import "testing"

func TestRound(t *testing.T) {
	t.Parallel()
	cases := []struct {
		f        float64
		expected float64
	}{
		{0.0, 0.0},
		{0.5, 1.0},
		{0.3, 0.0},
		{0.7, 1.0},
	}

	for _, c := range cases {
		rounded := Round(c.f)
		if rounded != c.expected {
			t.Errorf("Computing round(%v) = %v - expected %v", c.f, rounded, c.expected)
		}
	}
}

func TestRoundInt(t *testing.T) {
	t.Parallel()
	cases := []struct {
		f        float64
		expected int
	}{
		{0.0, 0},
		{0.5, 1},
		{0.3, 0},
		{0.7, 1},
	}

	for _, c := range cases {
		rounded := RoundInt(c.f)
		if rounded != c.expected {
			t.Errorf("Computing roundint(%v) = %v - expected %v", c.f, rounded, c.expected)
		}
	}
}

func TestRoundPrecision(t *testing.T) {
	t.Parallel()
	cases := []struct {
		f         float64
		precision int
		expected  float64
	}{
		{0.41353, 2, 0.41},
		{0.5111, 1, 0.5},
		{0.221, 0, 0},
		{0.722145553, 10, 0.722145553},
		{12.1423, -1, 10.0},
		{0, 4, 0},
	}

	for _, c := range cases {
		rounded := RoundPrecision(c.f, c.precision)
		if rounded != c.expected {
			t.Errorf("Computing round(%v) with precision %v = %v - expected %v", c.f, c.precision, rounded, c.expected)
		}
	}

}
