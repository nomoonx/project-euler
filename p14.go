package main

import (
	"fmt"
)

var (
	countMap = map[int]int{1: 1, 2: 2}
)

func main() {
	max := -1
	maxStart := 0
	for i := 3; i < 1000000; i++ {
		b := count(i)
		if b > max {
			max = b
			maxStart = i
		}
	}

	fmt.Println(max, maxStart)
}

func count(a int) int {
	if count, exists := countMap[a]; exists {
		return count
	}
	if a%2 == 0 {
		countMap[a] = 1 + count(a/2)
	} else {
		countMap[a] = 1 + count(a*3+1)
	}
	return countMap[a]
}
