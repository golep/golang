package main

import "fmt"

func main() {
	a := 20
	b := 30

	nama := "wowon"

	r1 := a == b

	fmt.Println(r1)
	fmt.Println(a >= b)
	fmt.Println(b >= a)
	fmt.Println(a != b)

	fmt.Println(12 >= 12)
	fmt.Println(13 >= 12)

	if a > b {
		println("ok")
	}

	if b > a {
		println("true")
	} else {
		println("false")
	}

	if nama == "wowok" {
		println("true")
	} else {
		println("false")
	}

	switch nama {
	case "wowok":
		println("1")

	case "wiwik":
		println("2")
	default:
		fmt.Println("masuk default")
	}

	switch true {
	case r1:
		fmt.Println("ok")
	default:
		fmt.Println("default")
	}

}
