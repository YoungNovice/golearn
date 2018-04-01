package main

import (
	"tree"
	"fmt"
)

// 扩展类型的方式
type myTreeNode struct {
 	node *tree.TreeNode
}

func (myNode *myTreeNode) postOder()  {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOder()
	right.postOder()
	fmt.Print(myNode.node.Value, " ")
}



// oop 只支持封装
func main() {
	var root tree.TreeNode
	// 直接创建一个结构
	root = tree.TreeNode{Value:3}
	// 左孩子是某个结果的指针
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{nil, nil, 5}
	root.Right.Left = new(tree.TreeNode)
	root.Left.Right = tree.CreateTreeNode(10)

	// 在slice 中创建treeNode 可以省略
	//nodes := []treeNode{ {value:3},{},{nil,nil, 6}}
	root.Traverse()

	fmt.Println("xxxxxxxx")
	myTreeNode := myTreeNode{&root}
	myTreeNode.postOder()


}

