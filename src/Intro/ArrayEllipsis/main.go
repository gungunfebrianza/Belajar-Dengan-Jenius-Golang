package main

import (
	"fmt"
)

type Book struct {
	bookID int32
	title  string
}

func main() {

	arr := [5]int64{1, 2, 3, 4, 5}

	fmt.Println("Panjang Array =", len(arr))

	for i := 0; i < len(arr); i++ {
		fmt.Println("Nilai i adalah", i)
	}
}
