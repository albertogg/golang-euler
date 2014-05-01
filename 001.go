package main

import (
	"fmt"
	"math"
)

func mod(x, y float64) float64 {
	var sum = float64(0)
	for i := float64(0); i < 1000; i++ {
		if (math.Mod(i, x) == 0) || (math.Mod(i, y) == 0) {
			sum += i
		}
	}
	return sum
}

func main() {
	fmt.Println(mod(3, 5))
}
