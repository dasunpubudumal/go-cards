package evenodd

import "fmt"

func EvenOrOdd(vals []int) {
	for _, val := range vals {
		if val%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("odd")
		}
	}
}
