package main

import "fmt"

func main() {
	for i := 997; i > 99; i-- {
		p := i*1000 + i%10*100 + i/100 + (i/10)%10*10
		for j := 999; j > 99; j-- {
			if p%j == 0 {
				k := p / j
				if k%1000 == k {
					fmt.Println(p, k, j)
					return
				}
			}
		}
	}
}
