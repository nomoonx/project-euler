package main

import "fmt"

func main() {
	result := []int{1}
	for i := 1; i <= 1000; i++ {
		plus := 0
		for j := len(result) - 1; j >= 0; j-- {
			result[j] = result[j]*2 + plus
			if result[j] >= 10 {
				plus = 1
				result[j] %= 10
			} else {
				plus = 0
			}
		}
		if plus > 0 {
			result = append([]int{1}, result...)
		}
	}
	sum := 0
	for _, num := range result {
		sum += num
	}
	fmt.Println(sum)
}
