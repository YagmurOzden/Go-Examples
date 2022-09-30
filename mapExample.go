package main

import (
	"fmt"
)

func main() {
	statePopulations := map[string]int{
		"Artvin":   12345672,
		"İstanbul": 98345656,
		"Ankara":   11293656,
		"İzmir":    28341237,
		"Edirne":   82189129,
		"Bursa":    92748103,
	}

	fmt.Println(statePopulations)
	statePopulations2 := make(map[string]int)
	statePopulations2 = map[string]int{
		"Artvin":   12345672,
		"İstanbul": 98345656,
		"Ankara":   11293656,
		"İzmir":    28341237,
		"Edirne":   82189129,
		"Bursa":    92748103,
	}
	fmt.Println(statePopulations2)
	delete(statePopulations2, "Ankara")
	fmt.Println(statePopulations2)

}
