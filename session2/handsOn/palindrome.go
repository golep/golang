package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	kata := bufio.NewScanner(os.Stdin)

	fmt.Println("input kata = ")
	kata.Scan()

	var balikan string = ""

	for i := len(kata.Text()) - 1; i >= 0; i-- {

		balikan = balikan + string(kata.Text()[i])
		fmt.Println(balikan)
	}

	if balikan == kata.Text() {
		fmt.Println("palindrom")
	} else {
		fmt.Println("bukan")
	}

}
