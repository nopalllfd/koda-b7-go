package calc

import "math"

const pi float32 = math.Pi

func Luas(r float32) (l float32) {
	return pi * r * r
}
func Keliling(r float32) (k float32) {
	return 2 * pi * r
}
func LuasDanKeliling(r float32) (float32, float32) {
	l := pi * r * r
	k := 2 * pi * r

	return l, k
}
