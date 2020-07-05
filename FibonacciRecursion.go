package main

import (
	"fmt"
	"time"
)

var start = time.Now()

// for 30 lenght total time spec 24.0636ms for 40 : 2.0045ms

func my_series(array []int64, s int) {
	a, b := 0, 1
	for i := 0; i <= s; i++ {
		//fmt.Println(a)
		array = append(array, int64(a))
		a, b = b, a+b
		//fmt.Println(i, a)
	}
	fmt.Println(array)
	end := time.Now()
	fmt.Println("Total time spend :", end.Sub(start))

}
func main() {
	fmt.Println(start)
	list_size := 50
	my_array := make([]int64, 0) //empty array creation, type int64 to get 100 approx
	my_series(my_array, list_size)
}
