package main

import (
	"fmt"
)

func main() {
	tree := [][]int{
		{75},
		{95, 64},
		{17, 47, 82},
		{18, 35, 87, 10},
		{20, 04, 82, 47, 65},
		{19, 01, 23, 75, 03, 34},
		{88, 02, 77, 73, 07, 63, 67},
		{99, 65, 04, 28, 06, 16, 70, 92},
		{41, 41, 26, 56, 83, 40, 80, 70, 33},
		{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
		{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
		{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
		{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
		{63, 66, 04, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
		{04, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23},
	}
	// tree := [][]int{
	// 	{3},
	// 	{7, 4},
	// 	{2, 4, 6},
	// 	{8, 5, 9, 3},
	// }
	sum := []int{tree[0][0]}
	for i, row := range tree {
		if i == 0 {
			continue
		}
		tempSum := make([]int, len(row))
		for j, element := range row {
			tempSum[j] = element
			if j == 0 {
				tempSum[j] += sum[j]
			} else if j == len(row)-1 {
				tempSum[j] += sum[j-1]
			} else {
				if sum[j] > sum[j-1] {
					tempSum[j] += sum[j]
				} else {
					tempSum[j] += sum[j-1]
				}
			}
		}
		sum = tempSum
		fmt.Println(sum)
	}
	max := -1
	for _, element := range sum {
		if element > max {
			max = element
		}
	}
	fmt.Println(max)
}
