package main

import (
	"fmt"
	"strconv"
)

func main() {

	s := "33.77773"

	if n, err := strconv.Atoi(s); err == nil {
		fmt.Println(n)
	} else {
		fmt.Println(err)
	}

}
