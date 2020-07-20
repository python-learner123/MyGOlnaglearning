package main

import (
	"fmt"
	"strconv"
)

var I int = 10 //global to lang
var i int = 20 //global to package

func main() {
	var i int = 20
	fmt.Println(i) //Local varible
	fmt.Printf("%v,%T\n", i, i)

	var j string = "20"
	j = strconv.Itoa(i)
	fmt.Printf("%v,%T\n", j, j) // with out 	"strconv" the result is ,string because it string is nothing but bytes....
	//asweer=20,string

	var k string = "hi ho"
	l := []byte(k)
	fmt.Printf("%v,%T\n", l, l) // result is byte string   [104 105 32 104 111],[]uint8

	a := 1 == 1

	fmt.Printf("%v,%T\n", a, a)

	z := 'a'
	fmt.Printf("%v,%T\n", z, z) // 97,int32  sinle corts is rune datatype which is int32
}
