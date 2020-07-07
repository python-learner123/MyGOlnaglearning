package main

import "fmt"

type Queue []interface{}

func (q *Queue) IsEmpty() bool {
	return (len(*q) == 0)
}

func (q *Queue) EnQueue(n interface{}) {

	*q = append(*q, n)
	fmt.Println("Element value", n, "Has been enQueued", *q)
}

func (q *Queue) DeQueue() {
	if q.IsEmpty() {
		fmt.Println("Queue is Empty")
	} else {
		value := (*q)[0]
		*q = (*q)[1:]
		fmt.Println("Element value", value, "Has been DeQueued", *q)
	}
}

func (q *Queue) Traverse() {
	fmt.Println("Items of Queues are--->")
	for k, z := range *q {
		fmt.Println("Key-->", k, "Value-->", z)
	}
}
func main() {
	var q Queue
	q.DeQueue() //q1 := &q    passing Q1 will also work
	for i := 1; i <= 10; i++ {
		q.EnQueue(i)
	}
	fmt.Println("BEFORE---->")
	q.Traverse()
	for i := 1; i <= 9; i++ {
		q.DeQueue()
	}
	fmt.Println("AFTER---->")
	q.Traverse()

}
