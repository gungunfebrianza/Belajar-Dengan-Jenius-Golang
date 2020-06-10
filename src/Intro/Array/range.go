package main

import (
	"fmt"
	"reflect"
)

func main() {

	numbers := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	kalimat := [2]string{"Hallo", "Maudy"}
	harga := [3]float64{7000.888, 999.555}

	for i, v := range numbers {
		fmt.Println("Array Element", i, "is", v)
	}

	for i, v := range kalimat {
		fmt.Println("Array String", i, "=", v)
	}

	for i, v := range harga {
		fmt.Println("Array Float64", i, "=", v)
	}

	hasil := 10 / 0.5

	fmt.Println(hasil)
	fmt.Println(reflect.TypeOf(hasil))
}
