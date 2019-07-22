package tree

import (
	"fmt"
)

type TreeNode struct {
	Value int
	Left, Right *TreeNode
}

func (node TreeNode) Print() {
	fmt.Println(node.Value)
}

func (node *TreeNode) SetValue(value int) {
	//need to pass the reference to change the value
	if node == nil {
		fmt.Println("setting value to node ignored")
		return 
	}
	node.Value = value
}

func (node *TreeNode) Traverse() {
	if node == nil {
		return 
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

func CreateNode(value int) *TreeNode {
	//return the address of the local area, save it on stack
	return &TreeNode{Value: value}
}
