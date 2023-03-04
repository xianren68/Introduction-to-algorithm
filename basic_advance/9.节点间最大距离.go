package basic_advance

import "math"

// 定义二叉树节点
type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

// 定义返回的数据
type Info struct {
	// 最大高度
	Height int
	// 最远距离
	Distance int
}

// 二叉树节点最大距离
func MaxDistance(head *TreeNode) int {
	return process(head).Distance
}
func process(head *TreeNode) *Info {
	if head == nil {
		return &Info{Height: 0, Distance: 0}
	}
	// 左右子树的返回值
	Left := process(head.Left)
	Right := process(head.Right)
	// 当前子树高度为它们中较大的+1
	height := int(math.Max(float64(Left.Height), float64(Right.Height)) + 1)
	// 当前子树的最远距离
	distance := int(math.Max(float64(Left.Height+Right.Height+1), math.Max(float64(Left.Distance), float64(Right.Distance))))
	return &Info{Height: height, Distance: distance}
}
