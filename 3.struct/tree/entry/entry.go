package main

import "tree"
import (
	"fmt"
)

type myTreeNode struct {
	node *tree.TreeNode
}

func (myNode *myTreeNode) posterOrder(){
	if myNode == nil || myNode.node == nil{
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.posterOrder()
	right := myTreeNode{myNode.node.Right}
	right.posterOrder()
	myNode.node.Print()
}

func defineTreeNodes() {
	var root tree.TreeNode
	root = tree.TreeNode{Value: 3}
	root.Left = &tree.TreeNode{Value: 4}
	root.Right = &tree.TreeNode{5, nil, nil}
	root.Right.Left = new(tree.TreeNode)
	root.Left.Right = tree.CreateNode(2)

	nodes := []tree.TreeNode {
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
	fmt.Println(root)
	root.Print()

	root.Right.Left.SetValue(4)
	root.Right.Left.Print()

	pRoot := &root
	pRoot.SetValue(300)
	pRoot.Print()
	pRoot.Traverse()
	
	fmt.Println()
	myRootNode := myTreeNode{&root}
	myRootNode.posterOrder()
}


func main() {
	defineTreeNodes()

}