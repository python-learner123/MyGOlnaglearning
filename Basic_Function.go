//Function syntex :==  func function_name(parameter_1 type, parameter_n type) return_type { //statement  }

package main

import "fmt"

func calculate(fi int, sec int) (int, int) {
	sum := fi + sec
	diff := fi - sec

	return sum, diff
}
func fizbuzz(num int) {
	if num%5 == 0 && num%3 == 0 {
		fmt.Println("FizBuzz")
	} else if num%5 == 0 {
		fmt.Println("Fiz")
	} else if num%3 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(num)
	}
}
func main() {

	x, y := 100, 50
	summ, diiff := calculate(x, y)
	fmt.Println("sum is :===", summ, "\n", "difference is :===", diiff)

	//Fiz Buzz test function
	for i := 1; i <= 30; i++ {
		fizbuzz(i)
	}
}
