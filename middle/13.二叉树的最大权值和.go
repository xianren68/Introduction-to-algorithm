package middle

import "math"

func MaxWeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 得到两边子树中权值更大的值+当前权值
	max := int(math.Max(float64(MaxWeight(root.LeftNode)), float64(MaxWeight(root.RightNode))))
	return max + root.Val
}
