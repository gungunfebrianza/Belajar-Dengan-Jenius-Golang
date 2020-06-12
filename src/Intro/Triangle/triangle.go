package main

import "fmt"

func main() {
	var rows, temp = 7, 1

	for i := 0; i < rows; i++ {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k <= i; k++ {
			if k == 0 || i == 0 {
				temp = 1
			} else {
				temp = temp * (i - k + 1) / k
			}
			fmt.Printf(" %d", temp)
		}
		fmt.Println("")
	}
}