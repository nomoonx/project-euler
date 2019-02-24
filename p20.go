package main

import "fmt"

func main() {
	result := []int{1}
	for i := 1; i <= 100; i++ {
		plus := 0
		for j := len(result) - 1; j >= 0; j-- {
			result[j] = result[j]*i + plus
			if result[j] >= 10 {
				plus = result[j] / 10
				result[j] %= 10
			} else {
				plus = 0
			}
		}
		if plus > 0 {
			for plus > 10 {
				result = append([]int{plus % 10}, result...)
				plus /= 10
			}
			result = append([]int{plus}, result...)

		}
		fmt.Println(result)
	}
	sum := 0
	for _, num := range result {
		sum += num
	}
	fmt.Println(sum)
}
