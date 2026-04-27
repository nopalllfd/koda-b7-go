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
		l := pi * r * r
		k := 2 * pi * r

		return l,k
	}

	func segitigaSikuSiku(h int){
		i := 1
		for i <= h{
			x := 0
			for x < i {
			fmt.Print("*")
			x++
			}
			i++
			fmt.Print("\n")
		}
	}

func main() {
	luas := luas(7)
	keliling := keliling(7)
	fmt.Printf("Luas: %.2f\n", luas)
	fmt.Printf("Keliling: %.2f\n", keliling)
	l, k := luasDanKeliling(7)
	fmt.Printf("Luas: %.2f Keliling: %.2f\n", l, k)
	segitigaSikuSiku(2)
}