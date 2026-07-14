package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	tinggi := bufio.NewScanner(os.Stdin)
	berat := bufio.NewScanner(os.Stdin)

	fmt.Println("input tinggi badan (cm) = ")
	tinggi.Scan()

	fmt.Println("input berat badan (kg)= ")
	berat.Scan()

	tinggiFloat, _ := strconv.ParseFloat(tinggi.Text(), 64)
	beratFloat, _ := strconv.ParseFloat(berat.Text(), 64)

	bmi := beratFloat / ((tinggiFloat * tinggiFloat) / 100)

	if bmi < 18.50 {
		fmt.Println(" Underweight")
	} else if bmi >= 18.50 && bmi <= 24.99 {
		fmt.Println(" healthy")
	} else if bmi >= 25 && bmi <= 29.99 {
		fmt.Println("overweight")
	} else if bmi >= 30 {
		fmt.Println("obese")
	} else {
		fmt.Println("out")
	}

}
