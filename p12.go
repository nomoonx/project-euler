package main

import "fmt"

func main() {
	i := 2
	for {

		a := i * (i + 1) / 2
		count := 2
		for j := 2; j < a; j++ {
			if a%j == 0 {
				count++
			}
		}
		// fmt.Println(count)
		if count > 500 {
			fmt.Println(a)
			break
		}
		if i%100 == 0 {
			fmt.Println(i, count)
		}
		i++
	}
}
