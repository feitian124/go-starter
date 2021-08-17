package mathx

import "math"

// RoundInt RoundWithFracInt For exact-value numbers, ROUND() uses the “round half up” rule:
// A value with a fractional part of .5 or greater is rounded up to the next integer if positive or down to the next
// integer if negative. (In other words, it is rounded away from zero.)
// A value with a fractional part less than .5 is rounded down to the next integer if positive or up to the next
// integer if negative. (In other words, it is rounded toward zero.)
func RoundInt(i int64, dec int) int64 {
	if dec >= 0 {
		return i
	}
	shift := math.Pow10(dec)
	r := math.Round(float64(i) * shift)
	for i := 1; i <= -dec; i++ {
		r *= 10
	}
	return int64(r)
}
