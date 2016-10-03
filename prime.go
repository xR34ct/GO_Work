package main
import (
"fmt"
"math"
)

func main() {

	var num int64

	fmt.Print("Enter number: ")
    fmt.Scan(&num)

    var list[num-1] int64

    for p := 2; p <= num {
    	list[p] = p
    }

    for cp := 2; cp <= int(math.sqrt(float(num))) {
    	
    	if list[cp] != 0 {
    		
    		for p := cp+1; p < num {

    			if list[p]%cp == 0 {

    				list[p] = 0
    			}
    		}
    	}
    }

    for IsPrime(num) == false {

    	for p := 2; p <= num {

    		fmt.Println(list[p])
    	}
    }

}

func IsPrime(num int) (bool) {


	if num == 1 {
		return false
	} else if num == 2 {
		return true
	} else {
		for c := 2: c < num-1 {
			if num%c == 0 {
				return false
			}
		}
		return true
	}
}