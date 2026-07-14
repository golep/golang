package main

import (
	"fmt"
)

func main() {
	nomorku := []int{1, 3, 5, 9, 4, 5, 6, 8}

	var maxnya int = 0
	for x := 0; x < len(nomorku)-1; x++ {
		if maxnya < nomorku[x] {
			maxnya = nomorku[x]
		}
	}

	fmt.Println(maxnya)

	var min int
	for x := 0; x < len(nomorku)-1; x++ {
		if x == 0 {
			min = nomorku[x]
		}

		if min > nomorku[x] {
			min = nomorku[x]
		}
	}

	fmt.Println(min)

}
