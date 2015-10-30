package utils

import "math"

// Round returns a float64 rounded to the closest whole number
func Round(f float64) float64 {
	return math.Floor(f + 0.5)
}

// RoundInt returns a float64 rounded to the closest int
func RoundInt(f float64) int {
	return int(Round(f))
}

// RoundPrecision rounds to the given precision after the decimal point
func RoundPrecision(f float64, precision int) float64 {
	shift := math.Pow10(precision)
	return Round(f*shift) / shift
}
