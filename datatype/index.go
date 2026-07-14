package main

import (
	"fmt"
)

func main() {

	is_married := false
	config := true

	fmt.Println(is_married, config)

	colors := [4]string{
		"red", "black", "green", "blue",
	}
	fmt.Println(colors)

	user := map[string]any{
		"name":  "Andi",
		"email": "andi@gmail.com",
		"city":  "Surabaya",
	}

	fmt.Println(user)
	fmt.Println(user["name"])

}
