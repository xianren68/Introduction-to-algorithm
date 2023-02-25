package binary_tree

import "math"

// 定义一个全局变量，存储每次的值
var preVal = math.MinInt

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if isValidBST(root.Left) {
		// 比大小
		if root.Val <= preVal {
			return false
		} else {
			preVal = root.Val
		}
	} else {
		return false
	}
	return isValidBST(root.Right)
}
