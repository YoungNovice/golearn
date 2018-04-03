package tree

import (
	"fmt"
)

// 定义树型节点
type TreeNode struct {
	Left, Right *TreeNode
	Value       int
}

// 接收者
func (node TreeNode) Print() {
	fmt.Print(node.Value)
}

func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	fmt.Print(node.Value, " ")
	node.Right.Traverse()
}

func (node *TreeNode) Traverse2() {
	node.TraverseFunc(func(node *TreeNode) {
		node.Print()
	})
	fmt.Println()
}

// 函数式编程实现二叉树便利 便利功能交给一个函数去实现
func (node *TreeNode) TraverseFunc(f func(*TreeNode)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

// 创建工厂函数
func CreateTreeNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}
