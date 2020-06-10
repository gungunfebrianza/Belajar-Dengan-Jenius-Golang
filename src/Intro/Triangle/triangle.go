package main

import (
	"fmt"
	"strings"
)

func main() {
	var rows = 10
	for i := 1; i <= rows; i++ {
		fmt.Println(strings.Repeat("na", rows))
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}

}
