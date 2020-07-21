package main

import (
	"fmt"
	"sync"
	"time"
)

//var mute sync.Mutex
var count_num = 0
var mute sync.Mutex

func count(n int) {

	for j := 1; j < 11; j++ {
		mute.Lock()
		//time.Sleep(time.Duration(rand.Int31n(2)) * time.Second)
		count_num = count_num + 1
		mute.Unlock()
	}
	// same as follow fmt.Println("total count in round", n, "is ", count_num)
	fmt.Println("total count in round", n, "is ", count_num)
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	for i := 1; i < 101; i++ {
		wg.Add(1)
		go func() {
			count(i)
			wg.Done()
		}()
	}
	wg.Wait()
	//time.Sleep(25 * time.Second)
	fmt.Println("closing main", count_num, "time taken by task: ", time.Now().Sub(start))

}

//without mutex closing main 810,closing main 910 etc
//with mutex closing main 1000 always
