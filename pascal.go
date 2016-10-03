package main
import (
    "fmt"
)

func main() {

	var row, place int64
	 
    fmt.Print("Enter row: ")
    fmt.Scan(&row)

    fmt.Print("Enter place: ")
    fmt.Scan(&place)

    fmt.Println(" ")

    for place > row { //If place > row loop
    	fmt.Println("Place cannot be bigger than row!")
    	fmt.Println(" ")
    	fmt.Println("Enter place: ")
    	fmt.Scan(&place)
    }

    fmt.Print("The number is: ")
    fmt.Println(PascalCalc(row,place))
}

func PascalCalc(r, p int64) (int64) { //Recursive function to calculate the number in Pascal's triangle

	if r == 1 || r == 2 || p == r || p == 1 {
		return 1
	} else {
			return (PascalCalc(r-1,p-1)+PascalCalc(r-1,p))
		}
}