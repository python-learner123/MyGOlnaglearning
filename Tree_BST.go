package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func (t *Tree) insert(i int) {
	if t.Root == nil {
		t.Root = &Node{Value: i}
	} else {
		t.Root.insert(i)
	}
}

func (n *Node) insert(i int) {
	if i <= n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: i}
		} else {
			n.Left.insert(i)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: i}
		} else {
			n.Right.insert(i)
		}
	}
}

func (n *Node) Preorder() {
	if n == nil {
		return
	} else {
		fmt.Print(n.Value, "-> ")
		n.Left.Preorder()
		n.Right.Preorder()
	}
}
func (n *Node) Postorder() {
	if n == nil {
		return
	} else {
		n.Left.Postorder()
		n.Right.Postorder()
		fmt.Print(n.Value, "-> ")
	}
}
func (n *Node) inorder() {
	if n == nil {
		return
	} else {
		n.Left.Postorder()
		fmt.Print(n.Value, "-> ")
		n.Right.Postorder()

	}
}
func main() {
	var t Tree
	t.insert(10)
	t.insert(5)
	t.insert(1)
	t.insert(8)
	t.insert(15)
	t.insert(12)
	t.insert(18)
	/*for i := 10; i <= 20; i++ {
		t.insert(i)
	}*/
	fmt.Println("Pre-Order ---->")
	t.Root.Preorder() // root first, then left  then right
	fmt.Println()
	fmt.Println("Post-Order ---->")
	t.Root.Postorder() //left first ,  then right then root
	fmt.Println()
	fmt.Println("In-Order ---->")
	t.Root.inorder() //left first , then root  then right
}
