package main

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

func main() {
	bt := &BinaryTree{}
	bt.Insert(100)
	bt.Insert(50)
	bt.Insert(120)
	bt.Insert(30)
	bt.Insert(200)
}
