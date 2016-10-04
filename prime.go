package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"time"
)

const file = "test.txt"

func main() {

	arg := flag.Int("num", 5, "number")
	flag.Parse()
	num := *arg

	/*var num int

	fmt.Print("Enter number: ")
	fmt.Scan(&num)*/

	list := make([]int, num-1)

	fmt.Println("Placing numbers!")
	for p := 1; p < num-1; p++ {
		list[p] = p
	}

	list_ := list

	CalcPrimes1(num, list)
	CalcPrimes2(num, list_)

	ListPrimes(num-1, list)
}

func IsPrime(num int) bool {

	if num == 1 {
		return false
	} else if num == 2 {
		return true
	} else {
		for c := 2; c < num-1; c++ {
			if num%c == 0 {
				return false
			}
		}
		return true
	}
}

func CalcPrimes1(num int, list []int) {
	start := time.Now()

	var sqrt int
	sqrt = int(math.Sqrt(float64(num)))
	//fmt.Println(sqrt)

	for cp := 2; cp <= sqrt; cp++ {
		if list[cp] != 0 {
			for p := cp + 1; p < num-1; p++ {
				if list[p]%(cp) == 0 {
					list[p] = 0
				}
			}
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Såll", elapsed)
}

func CalcPrimes2(num int, list []int) {
	start := time.Now()

	for cp := 2; cp < num-1; cp++ {
		if IsPrime(cp) == false {
			list[cp] = 0
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Ej såll", elapsed)
}

func ListPrimes(num int, list []int) {
	list[1] = 0
	var buffer bytes.Buffer
	P := 0

	fmt.Println("Printing!")
	write(os.O_CREATE|os.O_TRUNC|os.O_RDWR, strconv.Itoa(num+1))
	write(os.O_APPEND|os.O_RDWR, "\n")
	write(os.O_APPEND|os.O_RDWR, "\n")

	for p := 1; p < num-1; p++ {
		if list[p] != 0 {
			P++
			buffer.WriteString("Prime #")
			buffer.WriteString(strconv.Itoa(P))
			buffer.WriteString(" = ")
			buffer.WriteString(strconv.Itoa(list[p]))
			buffer.WriteString("\n")
			/*fmt.Print("Prime #")
			fmt.Print(P)
			fmt.Print(" = ")
			fmt.Println(list[p])*/
		}
	}
	fmt.Println(P)
	write(os.O_APPEND|os.O_RDWR, buffer.String())
}
func write(flag int, text string) {
	f, err := os.OpenFile(file, flag, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err := io.WriteString(f, text)
	if err != nil {
		fmt.Println(n, err)
		return
	}
	f.Close()
}
