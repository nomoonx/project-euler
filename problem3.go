package main

import (
	"fmt"
	"math"
)

func main() {
	upper := int64(math.Sqrt(600851475143))
	if upper%2 == 0 {
		upper++
	}
	for i := upper; i > 2; i -= 2 {
		if 600851475143%i == 0 && isPrime(i) {
			fmt.Println(i)
		}
	}

}

func isPrime(a int64) bool {
	upper := int64(math.Sqrt(float64(a)))
	if upper%2 == 0 {
		upper--
	}

	for i := upper; i > 2; i -= 2 {
		if a%i == 0 {
			return false
		}
	}
	return true
}
