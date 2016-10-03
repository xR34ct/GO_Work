package main
import (
    "fmt"
)

func main() {

	var row, place, number int
	 
    fmt.Print("Enter row: ")
    fmt.Scan(&row)

    fmt.Print("Enter place: ")
    fmt.Scan(&place)

    number = PascalCalc(row,place)

    fmt.Println("The number is: ")
    fmt.Println(number)

}

func PascalCalc(r int, p int) (int) {

	if r == 1 || r == 2 || p == r || p == 1 {

		return 1
	} else {
			return (PascalCalc(r-1,p-1)+PascalCalc(r-1,p))
		}
}