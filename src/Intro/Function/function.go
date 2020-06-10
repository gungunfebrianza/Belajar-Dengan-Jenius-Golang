package main

import (
	"fmt"
)

func addition(a int, b float64) int {
	return a + int(b)
}

func main() {
	fmt.Println(addition(2, 3.0))
}
