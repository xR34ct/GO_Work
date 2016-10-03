package main
import "fmt"
import "math"

func main() {

	var num int64

	fmt.Print("Enter number: ")
    fmt.Scan(&num)

    var list[num] int64

    for p := 1; p <= num; p++ {
    	list[p] = p
    }

    for cp := 1; cp <= int(math.sqrt(float(num))); cp++ {
    	
    	if list[cp] != 0 {
    		
    		for p := cp+1; p < num; p++ {

    			if list[p]%cp == 0 {

    				list[p] = 0
    			}
    		}
    	}
    }

    for IsPrime(num) == false {

    	for p := 1; p <= num; p++ {

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
		for c := 2: c < num-1; c++ {
			if num%c == 0 {
				return false
			}
		}
		return true
	}
}