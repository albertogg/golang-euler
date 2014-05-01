package main

import "fmt"

func main() {
	var a, b, sum, temp = 1, 2, 0, 0
	for a < 4000000 {
		if a%2 == 0 {
			sum += a
		}
		temp = b
		b = a + b
		a = temp
	}
	fmt.Println(sum)
}
