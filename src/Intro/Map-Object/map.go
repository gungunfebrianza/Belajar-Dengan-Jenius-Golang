package main

import (
	"fmt"
)

func main() {
	couple := make(map[string]int)
	couple["Meriam Beliana"] = 28
	couple["Bpk.Ariyanto"] = 51
	couple["Suzanna"] = 60
	fmt.Println(couple)
	delete(couple, "Suzanna")
	fmt.Println(couple)
	// map[Bpk.Ariyanto:51 Meriam Beliana:28 Suzanna:60]
  // map[Bpk.Ariyanto:51 Meriam Beliana:28]
}
