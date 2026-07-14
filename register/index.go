package main

import(
	"bufio"
	"fmt"
	"os"
)


func main(){
	fmt.Println("register")
	fmt.Println("============")


	name := bufio.NewScanner(os.Stdin)
	email := bufio.NewScanner(os.Stdin)

	fmt.Print("nama =")
	name.Scan()

	fmt.Print("emai =")
	email.Scan()


	fmt.Println("nama = "+ name.Text() + " email " +email.Text())



}