package main

import "fmt"

func main() {
	
	couple := make(map[string]int)


	couple["Maudy Ayunda"] = 28
	couple["Gun Gun Febrianza"] = 28
	couple["Afghan"] = 33

	couple = make(map[string]int)

	fmt.Println(len(couple)) //0
}