package main

//The recover function returns the value passed to panic function and has no side effects.
// That means if our goroutine is not panicking, recover function will return nil. So checking the return value of recover() to nil is a good way to know if your program is packing,

import "fmt"

func defFoo() {
	fmt.Println("defFoo() started")

	if r := recover(); r != nil {
		fmt.Println("WOHA! Program is panicking with value", r)
	}

	fmt.Println("defFoo() done")
}

func normMain() {

	fmt.Println("normMain() started")

	defer defFoo() // defer defFoo call

	panic("HELP") // panic here

	fmt.Println("normMain() done")
}

func main() {
	fmt.Println("main() started")

	normMain() // normal call

	fmt.Println("main() done")
}
