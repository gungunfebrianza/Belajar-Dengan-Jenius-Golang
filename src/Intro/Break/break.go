package main

import "fmt"

func main() {
	for index := 0; index < 5; index++ {
		if index == 4 {
			break
		}
		fmt.Println(index)
	}
}
