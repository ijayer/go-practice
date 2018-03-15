package main

import "fmt"

type Node struct {
	left  *Node
	data  interface{}
	right *Node
}

func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

func (n *Node) SetData(data interface{}) {
	n.data = data
}

func main() {
	root := NewNode(nil, nil)
	root.SetData("root-node")

	left := NewNode(nil, nil)
	left.SetData("left-node")

	right := NewNode(nil, nil)
	right.SetData("right-node")

	root.left = left
	root.right = right

	fmt.Printf("result: %v, %v, %v", root, root.left, root.right)
}
