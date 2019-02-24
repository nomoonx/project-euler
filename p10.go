package main

import "fmt"

func main() {
	primeMap := make(map[int]bool)
	boundry := 2000000

	sum := 2

	for i := 3; i <= boundry; i += 2 {
		if !primeMap[i] {
			for j := 3; j < boundry/i+1; j += 2 {
				primeMap[j*i] = true
			}
			sum += i
			// fmt.Println(sum, primeMap)
		}
	}
	fmt.Println(sum)
}
