package main

import "fmt"

//slice are actually the pointers to the elements of an array. It means if you change any element in a slice, the underlying array contents also will be affected.
func main() {
	var array [3]int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	fmt.Println(array)

	array2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println(array2)
	fmt.Println(len(array2)) //length
	fmt.Println(array2[1:3]) //slicing
	fmt.Println(array2[1])   // -1 will not work here

	array3 := array2[1:5]
	array3[0] = 10
	fmt.Println(array3)
	fmt.Println(append(array3, 1000))      //append variable in array
	fmt.Println(append(array3, array3...)) //append slice to slice not possible, hence second slice needs to be unpacked as values to append

}

/*
package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	a = append(a, 8, 9)
	fmt.Print(len(a), cap(a)) //if more then the capacity after append the go will double the capacity
}

*/
//https://medium.com/rungo/the-anatomy-of-slices-in-go-6450e3bb2b94
// cpoy will copy the data till the length id copy from array with 10 data to null array then nothing will copy
//similary if copy data from 10 lenght array to 20 length arry then only first 10 will copy

// s1 := make([]int ,2,2) create a slice of len 2 with all data will be zero
//s2:= []int{1,2} create a slice of len 2 with data 1 and 2
// s3:= []int  its a slice refernce with no data mean the len will be zero ,hence  by copy command nothing will update
//[start:end] extract operator, end will not include if start is given like [2:4] but included in [:4]
