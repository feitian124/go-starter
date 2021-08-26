package mathx

import (
	"fmt"
	"math"
)

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

	// 取整数部分, 12345 -> 12300
	shift := int64(math.Pow10(-dec))
	x := i / shift * shift
	fmt.Printf("%+v", x)

	// 取剩余部分, 12345 -> 45
	y := i - x

	// 判断剩余部分第一位是否大于 5, 4 < 5
	xx := y * 10 / shift
	fmt.Printf("%+v", xx)
	if xx >= 5 || x <= -5 {
		x += shift
	}

	return x
}

// Transcribed from Postgres' src/port/rint.c, with c-style comments preserved
// for ease of mapping.
//
// https://github.com/postgres/postgres/blob/REL9_6_3/src/port/rint.c
func round(x float64) float64 {
	/* Per POSIX, NaNs must be returned unchanged. */
	if math.IsNaN(x) {
		return x
	}

	/* Both positive and negative zero should be returned unchanged. */
	if x == 0.0 {
		return x
	}

	roundFn := math.Ceil
	if math.Signbit(x) {
		roundFn = math.Floor
	}

	/*
	 * Subtracting 0.5 from a number very close to -0.5 can round to
	 * exactly -1.0, producing incorrect results, so we take the opposite
	 * approach: add 0.5 to the negative number, so that it goes closer to
	 * zero (or at most to +0.5, which is dealt with next), avoiding the
	 * precision issue.
	 */
	xOrig := x
	x -= math.Copysign(0.5, x)

	/*
	 * Be careful to return minus zero when input+0.5 >= 0, as that's what
	 * rint() should return with negative input.
	 */
	if x == 0 || math.Signbit(x) != math.Signbit(xOrig) {
		return math.Copysign(0.0, xOrig)
	}

	/*
	 * For very big numbers the input may have no decimals.  That case is
	 * detected by testing x+0.5 == x+1.0; if that happens, the input is
	 * returned unchanged.  This also covers the case of minus infinity.
	 */
	if x == xOrig-math.Copysign(1.0, x) {
		return xOrig
	}

	/* Otherwise produce a rounded estimate. */
	r := roundFn(x)

	/*
	 * If the rounding did not produce exactly input+0.5 then we're done.
	 */
	if r != x {
		return r
	}

	/*
	 * The original fractional part was exactly 0.5 (since
	 * floor(input+0.5) == input+0.5).  We need to round to nearest even.
	 * Dividing input+0.5 by 2, taking the floor and multiplying by 2
	 * yields the closest even number.  This part assumes that division by
	 * 2 is exact, which should be OK because underflow is impossible
	 * here: x is an integer.
	 */
	return roundFn(x*0.5) * 2.0
}
