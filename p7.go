package main

import "fmt"

func main() {
	primeMap := make(map[int]bool)
	boundry := 200000

	count := 0

	for i := 2; i <= boundry; i++ {
		if !primeMap[i] {
			for j := 2; j < boundry/i; j++ {
				primeMap[j*i] = true
			}
			count++
		}
		if count == 10001 {
			fmt.Println(count, i)
			return
		}
	}
	fmt.Println(count)
}
