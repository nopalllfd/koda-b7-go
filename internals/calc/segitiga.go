package calc

import "fmt"

func SegitigaSikuSiku(h int) {
	i := 1
	for i <= h {
		x := 0
		for x < i {
			fmt.Print("*")
			x++
		}
		i++
		fmt.Print("\n")
	}
}
