package main

import "fmt"

type Stack []interface{}  // to support all types of datatype

func (s *Stack) IsEmpty() bool {
	return (len(*s) == 0)
}

func (s *Stack) Push(item interface{}) {
	fmt.Println("Item ", item, " has been Pushed in Stack")
	*s = append(*s, item)
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		fmt.Println("Stack is Empty")
	}
	stack_index := len(*s) - 1     // index starts from zero
	last_item := (*s)[stack_index] // without ()   ->  invalid operation: s[stack_index] (type *Stack does not support indexing)
	*s = (*s)[:stack_index]        // it excludes the last in range
	fmt.Println("Item ", last_item, " has been Poped from Stack")

}
func (s *Stack) Traverse() {
	fmt.Println("-------Items of Stack-------------")
	for _, z := range *s {
		fmt.Println(z)
	}
	fmt.Println("----------------------------------")
}
func main() {

	var item int = 10
	var pull int = 5
	var stack Stack
	stack.Push("String")
	stack.Push(3.14)
	for i := 1; i <= item; i++ {
		stack.Push(i)
	}
	fmt.Println("BEFORE--------->")
	stack.Traverse()
	for i := 0; i <= pull; i++ {
		stack.Pop()
	}
	fmt.Println("AFTER--------->")
	stack.Traverse()
}
