package main
import "fmt"
import "math"

func main() {

	var num int

	fmt.Print("Enter number: ")
    fmt.Scan(&num)

    list := make([]int, num-1)

    for p := 1; p < num-1; p++ {
    	list[p] = p
    }



    for cp := 2; cp <= int(math.Sqrt(float64(num))); cp++ {
    	if list[cp] != 0 {
    		for p := cp+1; p < num-1; p++ {
    			if list[p]%(cp) == 0 {
    				list[p] = 0
    			}
    		}
    	}
    }

    list[1] = 0

    P := 1
    for p := 1; p < num-1; p++ {
    	if list[p] != 0 {
    		fmt.Print("Prime #")
    		fmt.Print(P)
    		fmt.Print(" = ")
    		fmt.Println(list[p])
    		P++
    	}
    }

}

func IsPrime(num int) (bool) {


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