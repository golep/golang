package main

import (
	"fmt"
)

func main() {
	nomorku := []int{2, 4, 1, 9, 7, 8, 20}
	max := nomorku[0]
	for x := 0; x < len(nomorku); x++ {
		for y := x + 1; y < len(nomorku); y++ {
			if max < nomorku[y]-nomorku[x] {
				max = nomorku[y] - nomorku[x]
			}
		}
	}
	fmt.Println(max)

	zxc := 5
	nomorku2 := [][]int{}
	current := 1
	for x := 0; x < zxc; x++ {
		temp := []int{}
		for y := 0; y < zxc; y++ {
			temp = append(temp, current)
			current++
		}
		nomorku2 = append(nomorku2, temp)
	}
	fmt.Println(nomorku2)

}
