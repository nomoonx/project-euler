package main

import "fmt"

func main() {
	count := [21][21]int{}
	for i := 0; i < 21; i++ {
		count[i][0] = 1
		count[0][i] = 1
	}
	for i := 1; i < 21; i++ {
		for j := 1; j < 21; j++ {
			count[i][j] = count[i-1][j] + count[i][j-1]
		}
	}
	fmt.Println(count[20][20])

}
