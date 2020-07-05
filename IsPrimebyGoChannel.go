package main

//10000  records   in 261301 microsecond
import (
	"fmt"
	"time"
)

// Testing
// print only prime go

func getprime(rangp int, chp chan int) {
	for i := 2; i <= rangp; i++ {
		//fmt.Println("value of I=", i)
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
			chp <- i
		}
	}
	close(chp)

}
func main() {
	start := time.Now()
	ch1 := make(chan int)
	rang := 10000
	go getprime(rang, ch1)

	/*for range ch1 {
		time.Sleep(10 * time.Millisecond)
		fmt.Println("PRIME----------->", <-ch1)
	}*/
	//time.Sleep(10 * time.Millisecond)
	for {
		//fetch data from channel
		x, flag := <-ch1

		//flag is true if data is received from the channel
		//flag is false when the channel is closed
		if flag == true {
			fmt.Println(x)
		} else {
			//fmt.Println("Empty channel")
			break
		}
	}
	fmt.Println("Total time took to complete=", time.Now().Sub(start).Microseconds())
}
