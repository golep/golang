package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	usd := bufio.NewScanner(os.Stdin)
	rupiah := bufio.NewScanner(os.Stdin)


	fmt.Print("kurs usd idr =")
	usd.Scan()

	fmt.Print("berapa usd mu =")
	rupiah.Scan()

	kursUsd,_ := strconv.ParseFloat(usd.Text(), 64)
	jumlahUSD,_ := strconv.ParseFloat(rupiah.Text(), 64)

	uangmu := kursUsd * jumlahUSD

	fmt.Printf("Uang kamu adalah = Rp %.2f\n", uangmu)
	fmt.
}