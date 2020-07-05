package main

import "fmt"

func check_pallindrom(num int) {
	number, temp := num, 0
	//s := make([]int, 10)
	for {
		digit := num % 10
		temp = temp*10 + digit
		num = num / 10
		if num == 0 {
			break
		}
	}
	if number == temp {
		fmt.Println(number, " is a pallindrom number")
		//s = append(s, number)
	}
}

func main() {
	var i int
	//fmt.Scanf("%d", &i)
	//i = 101
	//fmt.Println(i)
	for i = 1001; i <= 9999; i++ {
		check_pallindrom(i)
	}

}
