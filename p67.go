package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	workingDirectory, _ := os.Getwd()
	filepath := workingDirectory + "/p067_triangle.txt"

	file, _ := os.Open(filepath)
	defer file.Close()

	reader := bufio.NewReader(file)

	tree := [][]int{}

L:
	for {
		line, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		for isPrefix {
			var tempLine []byte
			tempLine, isPrefix, err = reader.ReadLine()
			if err == io.EOF {
				break L
			}
			line = append(line, tempLine...)
		}
		params := strings.Split(string(line), " ")
		newRow := make([]int, len(params))
		for i, param := range params {
			newRow[i], _ = strconv.Atoi(param)
		}
		tree = append(tree, newRow)
	}

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
