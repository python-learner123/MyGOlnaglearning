package main

import (
	"fmt"
	"time"
)

var start = time.Now()

func count(k string) {
	//var mydict map[string]int //panic: assignment to entry in nil map because the map is just created as refernce but no underlying datastruncture
	mydict := make(map[string]int) // it will create underlying DS for map
	for _, key := range k {
		//fmt.Println(string(key))
		_, ok := mydict[string(key)]
		if ok == true {
			mydict[string(key)] = mydict[string(key)] + 1
		} else {
			mydict[string(key)] = 1
		}
	}
	t := 1
	for keyy, valuee := range mydict {
		if valuee == 1 {
			fmt.Println("Key", keyy, "is unique character")
			t = 0
		} else {
			fmt.Println("Key", keyy, "Repeated", valuee, "times ")
		}
	}
	if t == 1 {
		fmt.Println("No unique character")
	}
	end := time.Now()
	fmt.Println("Total time spend :", end.Sub(start))
}

func main() {
	fmt.Println("Enter the String for which character needs to count: ")
	var key string
	//fmt.Scanln(&key)
	key = "aabbccdkjhwqieyuwqoudshlaj;xajxwqlkbiu21epqie[p23k;lkslkanbckjadn;wjd;qwjepo321hjelqkwjelqwk"
	fmt.Println(key)
	count(key)
}
