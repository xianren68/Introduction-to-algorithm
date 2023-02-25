package binary_tree

import "math"

func isBalanced(root *TreeNode) bool {
	isTrue, _ := rec(root)
	return isTrue
}

// 递归的方法，返回每棵子树的深度和它是不是满足平衡二叉树
func rec(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	lbl, lde := rec(root.Left)
	rbl, rde := rec(root.Right)
	// 左右都为平衡且深度差不超过1，则自身也为平衡
	// 返回true,且深度为最大值+1
	if (lbl && rbl) && (math.Abs(float64(lde-rde)) <= 1) {
		return true, int(math.Max(float64(lde), float64(rde))) + 1
	}
	return false, 0
}
