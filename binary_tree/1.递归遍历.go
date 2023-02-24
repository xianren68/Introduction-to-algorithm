package binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 先序
func RecPreTraversal(root *TreeNode) []int {
	// 创建返回的数组
	res := make([]int, 0)
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	rel := RecPreTraversal(root.Left)
	res = append(res, rel...)
	rer := RecPreTraversal(root.Right)
	res = append(res, rer...)
	return res
}

// 中序
func RecInTraversal(root *TreeNode) []int {
	// 创建返回的数组
	res := make([]int, 0)
	if root == nil {
		return res
	}

	rel := RecInTraversal(root.Left)
	res = append(res, rel...)
	res = append(res, root.Val)
	rer := RecInTraversal(root.Right)
	res = append(res, rer...)
	return res
}

// 后序
func RecPostTraversal(root *TreeNode) []int {
	// 创建返回的数组
	res := make([]int, 0)
	if root == nil {
		return res
	}

	rel := RecPostTraversal(root.Left)
	res = append(res, rel...)
	rer := RecPostTraversal(root.Right)
	res = append(res, rer...)
	res = append(res, root.Val)
	return res
}
