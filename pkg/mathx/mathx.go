package mathx

import "math"

// RoundInt rounds the argument i to dec decimal places using "round half up" rule.
// dec defaults to 0 if not specified. dec can be negative
func RoundInt(i int64, dec int) int64 {
	// is itself when dec >= 0
	if dec >= 0 {
		return i
	}
	// The integer part of a fraction
	shift := math.Pow10(dec)
	intPart := math.Round(float64(i) * shift)
	r := int64(intPart)
	// the zero part
	for i := 1; i <= -dec; i++ {
		println("test hook, ming, r=", r, "i=", i, ", -dec=", -dec)
		r *= 10
	}
	return r
}
