package main

import (
	"fmt"
	"strings"
)

func main() {
	bank_name := "Bank Jatim"

	rep := strings.Replace(bank_name, "J", "Y", 1)
	contains := strings.Contains(bank_name, "Bank")

	fmt.Println(rep, contains)

}
