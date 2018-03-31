package tree

import (
	"fmt"
)

// 定义树型节点
type TreeNode struct {
	Left, Right *TreeNode
	Value int
}

// 接收者
func (node TreeNode) Print()  {
	fmt.Print(node.Value)
}

func (node *TreeNode) Traverse()  {
	if node == nil {
		return
	}
	node.Left.Traverse()
	fmt.Print(node.Value, " ")
	node.Right.Traverse()
}


// 创建工厂函数
func CreateTreeNode(value int) *TreeNode {
	return &TreeNode{Value:value}
}

