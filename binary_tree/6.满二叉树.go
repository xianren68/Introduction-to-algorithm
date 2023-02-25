package binary_tree

func isFBT(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	if root.Left != nil && root.Right != nil {
		return isFBT(root.Left) && isFBT(root.Right)
	}
	// 两个节点不全
	return false
}
