/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

package main

import "fmt"

func main() {
	var n = 600851475143
	s := make([]int, 1)
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			s[0] = i
			n = n / i
		}
	}
	fmt.Printf("largest prime factor %d", s[0])
}
