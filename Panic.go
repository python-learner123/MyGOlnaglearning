package main

import "fmt"

/*
A panic is similar to an exception which can occur at runtime.\
 A panic can be thrown by the runtime or can be deliberately thrown by the programmer to handle the worst case scenario.
*/
func arry(arr []int, i int) int {
	if len(arr) > i {
		return arr[i]
	} else {
		panic("array is out of boud")
	}
}
func currvlaue(arr []int) {
	fmt.Println(arr)
}

func main() {
	array := []int{1, 2, 3} //slice
	defer currvlaue(array)
	fmt.Println(arry(array, 1))
	//panic
	fmt.Println(arry(array, 4))
}
