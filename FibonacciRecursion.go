package main

import (
	"fmt"
	"time"
)

var start = time.Now()

// for 30 lenght total time spec 24.0636ms , for 40 2.9959799s
//start := time.Now()  syntax error: non-declaration statement outside function body
//:= only works in functions and the lower case 't' is so that it is only visible to the package
//0 1 1 2 3 5 8 13

func recorsion(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return (recorsion(n-1) + recorsion(n-2))
	}
}

func my_series(array []int, s int) {
	for i := 0; i <= s; i++ {
		value := recorsion(i)
		//fmt.Println(i, value)
		array = append(array, value)
	}
	fmt.Println(array)
	end := time.Now()
	fmt.Println("Total time spend :", end.Sub(start))

}

func main() {
	fmt.Println(start)
	list_size := 40
	my_array := make([]int, 0) //empty array creation
	my_series(my_array, list_size)
}
