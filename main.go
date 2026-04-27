package main

import (
	"fmt"
	"math"
)

const pi float64 = math.Pi
	func luas(r float64) (l float32)  {
		return float32(pi * r * r)
	}
	func keliling(r float64)(k  float32){
		return float32(2 * pi * r)
	}
	func luasDanKeliling(r float64) (float32, float32){
		l := luas(r)
		k := keliling(r)

		return l,k
	}

func main() {
	luas, keliling := luasDanKeliling(7)
	fmt.Printf("Luas: %.2f Keliling: %.2f", luas, keliling)
}