package main

import "fmt"
// ------------ ByPython Learner for GoLang on 05-JUL-20
type Numbers struct {
	Value    int
	Valuestr string
	Next     *Numbers
}
type mylist struct {
	Mylistname string
	Head       *Numbers
}

func createlist(name string) *mylist {
	return &mylist{
		Mylistname: name,
	}
}

func (l *mylist) addnum(value int, valuestr string) {
	numbers := &Numbers{
		Value:    value,
		Valuestr: valuestr,
	}
	if l.Head == nil {
		l.Head = numbers
	} else {
		currentnode := l.Head
		//test := 1
		for currentnode.Next != nil {
			currentnode = currentnode.Next
			/*test = test + 1
			fmt.Print(test)
			if test == 10 {
				break
			}*/

		}
		currentnode.Next = numbers
	}
}

func (l *mylist) traverselist() {
	if l.Head == nil {
		fmt.Println("Empty list")
	} else {
		node := l.Head
		//fmt.Printf(, *node)
		fmt.Println(node.Value, node.Valuestr)
		for node.Next != nil {
			node = node.Next
			fmt.Println(node.Value, node.Valuestr)
		}
	}
}

func (l *mylist) delnum(value int) {
	fmt.Println("Value==============", value)
	if l.Head == nil { //if head is null
		fmt.Println("Empty list")
	}

	node := l.Head
	if node.Value == value {
		fmt.Println("Deleting the Head list")
		l.Head = l.Head.Next //if value is in head
	}
	//test := 1
	for node.Next != nil { //if value in node
		if node.Next.Value == value {
			if node.Next.Next == nil { //condition if value in last node
				fmt.Println("Deleting the last element from list")
				node.Next = nil
				break
			}
			fmt.Println("Intermediate element deletion")
			node.Next = node.Next.Next
		}
		node = node.Next

	}
}
func main() {
	lisname := "MyFirstlinklist"
	linklistname := createlist(lisname)
	fmt.Println(linklistname)
	linklistname.addnum(1, "One")
	linklistname.addnum(2, "Two")
	linklistname.addnum(3, "Three")
	for i := 4; i <= 10; i++ {
		linklistname.addnum(i, "BY Loop")
	}
	linklistname.traverselist()
	linklistname.delnum(1)
	linklistname.delnum(5)
	linklistname.delnum(10)
	fmt.Println("AFTER----------->")
	linklistname.traverselist()
}
