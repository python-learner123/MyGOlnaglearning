// go do not allow to create a named function inside of other named function
/*Go functions may be closures. A closure is a function value that references variables from outside its body.
The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.*/

package main

import "fmt"

func namedfunc() func() {
	var b = 20
	fmt.Println(" Hello from named function with value:", b)
	return func() {
		fmt.Println(" Hello from anonymous function with value:", b) // garbage collecter do not clear it though the scope is over,
		// because anonymous fucntion having this value for use
	}
}

func main() {
	var a = 10
	fmt.Println(" Hello from main function with value:", a)
	f := namedfunc()
	// peform many operaiton here before calling anonymous fucntion
	// it will always keep value from named fuction untill we use it
	fmt.Println(" Hello from main function again with value:", a)
	f()
}
