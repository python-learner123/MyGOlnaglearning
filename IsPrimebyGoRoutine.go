package main

//10000  records   in 26926 microsecond

import (
	"fmt"
	"sync"
	"time"
)

// Testing
// print only prime go

func getprime(i int) {
	isprime := true
	for j := 2; j < i; j++ {
		//fmt.Println("value of J=", j)
		if i%j == 0 {
			isprime = false
			break
		}
	}
	if isprime == true {
		//fmt.Println("value of PRIME INSIDE", i)
		fmt.Println(i)
	}
}
func main() {
	start := time.Now()
	var vg sync.WaitGroup
	rang := 10000
	for r := 2; r <= rang; r++ {
		//fmt.Println("Here are the go routine number=", r)
		vg.Add(1)
		func() {
			go getprime(r)
			vg.Done()
		}()

	}
	vg.Wait()

	fmt.Println("Total time took to complete=", time.Now().Sub(start).Microseconds())
}
