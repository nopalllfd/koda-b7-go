package main

import (
	"fmt"
	"math"
)

	const pi float32 = math.Pi
	func luas(r float32) (l float32)  {
		return pi * r * r
	}
	func keliling(r float32)(k  float32){
		return 2 * pi * r
	}
	func luasDanKeliling(r float32) (float32, float32){
		l := luas(r)
		k := keliling(r)

		return l,k
	}

func main() {
	luas, keliling := luasDanKeliling(7)
	fmt.Printf("Luas: %.2f Keliling: %.2f", luas, keliling)
}