/*
Variadic functions are also functions but they can take an infinite or variable number of arguments.
we have seen this in slices lesson when append function accepted a variable number of arguments.
func f(elem ...Type)
A typical syntax of a variadic function looks like above. ... operator called as pack operator instructs Go to store all arguments of type Type in elem parameter.

As from append function syntax, we can’t pass another slice as an argument, it has to be one or many arguments of type Type. Hence, instead, we will use the unpack operator ... to unpack slice into the series of arguments (which is acceptable by append function).
append(s1, s2...)

*/

package main

import "fmt"

func varidicfunc(factor int, s ...int) []int {
	for index, value := range s {
		s[index] = factor + value
	}
	return s
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := varidicfunc(3, slice...) // unpacking the slice in order to pass as validic argument
	fmt.Println(result)
	result2 := varidicfunc(2, 9, 8, 7, 6, 5)
	fmt.Print(result2)
}
