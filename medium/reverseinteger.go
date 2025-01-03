package medium

import "math"

const step int = 10

func reverse(x int) int {
	if -2<<31 > x || x > 2<<31 {
		return 0
	}
	xabs := int(math.Abs(float64(x)))
	numcount := int(math.Floor(math.Log10(float64(xabs)))) + 1
	coaf := int(math.Pow10(numcount))
	iterateStep := 1
	var res int
	var n int
	for i := numcount; i > 0; i-- {
		n++
		val := xabs % int(math.Pow10(n)) / iterateStep
		if val > 0 {
			res += (coaf / step) * val
		}
		coaf /= step
		iterateStep *= step
	}
	if x < 0 {
		res *= -1
	}
	return res
}
