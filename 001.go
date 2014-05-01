/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we
get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

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
