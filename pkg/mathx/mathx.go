package mathx

import "math"

// RoundInt rounds the argument i to dec decimal places using "round half up" rule.
// dec defaults to 0 if not specified. dec can be negative
func RoundInt(i int64, dec int) int64 {
	// is itself when dec >= 0
	if dec >= 0 {
		return i
	}
	//any digits in excess of 30 (or -30) are truncated as per mysql manual
	//if dec < -30 {
	//	dec = -30
	//}
	shift := math.Pow10(-dec)
	intPart := math.Round(float64(i) / shift)
	return int64(intPart) * int64(shift)
}
