package main

import (
	"fmt"
)

// 定义树型节点
type treeNode struct {
	left, right *treeNode
	value int
}

// 接收者
func (node treeNode) print()  {
	fmt.Print(node.value)
}

func (node *treeNode) traverse()  {
	if node == nil {
		return
	}
	node.left.traverse()
	fmt.Print(node.value, " ")
	node.right.traverse()
}


// 创建工厂函数
func createTreeNode(value int) *treeNode  {
	return &treeNode{value:value}
}

// oop 只支持封装
func main() {
	var root treeNode
	// 直接创建一个结构
	root = treeNode{value:3}
	// 左孩子是某个结果的指针
	root.left = &treeNode{}
	root.right = &treeNode{nil, nil, 5}
	root.right.left = new(treeNode)
	root.left.right = createTreeNode(10)
	
	

	// 在slice 中创建treeNode 可以省略
	//nodes := []treeNode{ {value:3},{},{nil,nil, 6}}
	root.traverse()

}
