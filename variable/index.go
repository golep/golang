package main

import(

	"fmt"

)


func main(){
	name := "golep"

	fmt.Println(name[0],string(name[0]))

	name = "halo "+name

	fmt.Println(name)

	var num1 = 100
	num2:=150

	add:= num1 + num2
	substract := num2 - num1
	multiply := num1*num2
	devide := num2/num1

	fmt.Println(add,substract,multiply,devide)
}