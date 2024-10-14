package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	root *Node
}

func (bt *BinaryTree) Insert(value int) {
	bt.root = InsertNode(bt.root, value)
}

func InsertNode(node *Node, value int) *Node {
	if node == nil {
		return &Node{value, nil, nil}
	}

	if value < node.Value {
		node.Left = InsertNode(node.Left, value)
	}
	if value > node.Value {
		node.Right = InsertNode(node.Right, value)
	}
	return node
}

func (bt *BinaryTree) DFS() {
	DFSHelper(bt.root)
}

func DFSHelper(node *Node) {
	if node == nil {
		return
	}
	DFSHelper(node.Left)
	fmt.Printf("%d ", node.Value)
	DFSHelper(node.Right)
}

func (bt *BinaryTree) BFS() {
	if bt.root == nil {
		return
	}

	queue := []*Node{bt.root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Printf("%d ", node.Value)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

func main() {
	bt := &BinaryTree{}
	bt.Insert(100)
	bt.Insert(50)
	bt.Insert(120)
	bt.Insert(30)
	bt.Insert(200)

	fmt.Println("DFS")
	bt.DFS()

	fmt.Println()

	fmt.Println("BFS")
	bt.BFS()
}
