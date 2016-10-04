package main

import (
	"fmt"
	"time"
)

func main() {

	var row, place int

	fmt.Print("Enter row: ")
	fmt.Scan(&row)

	for row < 1 {
		fmt.Println("Row cannot be lesser than 0!")
		fmt.Println(" ")
		fmt.Print("Enter row: ")
		fmt.Scan(&place)
	}

	fmt.Print("Enter place: ")
	fmt.Scan(&place)

	for place < 0 {
		fmt.Println("Place cannot be lesser than 0!")
		fmt.Println(" ")
		fmt.Print("Enter place: ")
		fmt.Scan(&place)
	}

	fmt.Println(" ")

	for place > row { //If place > row loop
		fmt.Println("Place cannot be bigger than row!")
		fmt.Println(" ")
		fmt.Print("Enter place: ")
		fmt.Scan(&place)
	}

	fmt.Print("The number is: ")
	start := time.Now()
	fmt.Println(PascalCalc(row, place))
	elapsed := time.Since(start)
	fmt.Println("Såll", elapsed)
}

func PascalCalc(r, p int) int64 { //Recursive function to calculate the number in Pascal's triangle

	if r == 1 || r == 2 || p == r || p == 1 {
		return 1
	} else {
		return (PascalCalc(r-1, p-1) + PascalCalc(r-1, p))
	}
}
