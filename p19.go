package main

import "fmt"

var (
	monthDayArray = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
)

func main() {
	count := 0
	startday := 2
	for year := 1901; year <= 2000; year++ {
		for index, day := range monthDayArray {
			if year%4 == 0 && index == 2 && year%400 > 0 {
				day++
			}
			startday += (day + 1) % 7
			startday %= 7
			// fmt.Println(startday)
			if startday == 0 && (index != 11 || year != 2000) {
				fmt.Println(year, index)
				count++
			}
		}
	}
	fmt.Println(count)
}
