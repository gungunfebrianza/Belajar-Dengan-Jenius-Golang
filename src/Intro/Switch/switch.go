package main

import (
	"fmt"
	"time"
)

func main() {

	dayBorn := time.Sunday

	switch dayBorn {
	case time.Monday:
		fmt.Println("This is Monday!")
	case time.Tuesday:
		fmt.Println("This is Tuesday")
	case time.Wednesday:
		fmt.Println("This is Wednesday")
	case time.Thursday:
		fmt.Println("This is Thursday")
	case time.Friday:
		fmt.Println("This is Friday")
	case time.Sunday:
		fmt.Println("Sekarang Hari Minggu")
	default:
		fmt.Println("Day, is not valid! Error!")
	}
}
