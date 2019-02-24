package main

import "fmt"

const (
	upperBound = 10000
)

var (
	amicableMap = make(map[int]int)
)

func main() {
	for num := 1; num < upperBound/2; num++ {
		for j := 2; j <= upperBound/num; j++ {
			amicableMap[j*num] += num
		}
	}
	fmt.Println(amicableMap[220])
	sum := 0
	for num := 1; num < upperBound; num++ {
		if amicableMap[num] < num {
			if num == amicableMap[amicableMap[num]] {
				fmt.Println(num, amicableMap[num])
				sum += num + amicableMap[num]
			}
		}
	}
	fmt.Println(sum)
}
