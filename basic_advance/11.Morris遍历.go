package basic_advance

func MorrisPre(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	// 定义返回的数组
	res := make([]int, 0)
	// 当前节点
	cur := root
	for cur != nil {
		// 存在左节点
		if cur.Left != nil {
			// 左节点最右的值
			var LeftRight *TreeNode = cur.Left
			// 找到最右节点
			for LeftRight.Right != nil && LeftRight.Right != cur {
				LeftRight = LeftRight.Right
			}
			// 判断是第几次来到当前节点
			if LeftRight.Right == nil { // 第一次到达
				res = append(res, cur.Val)
				LeftRight.Right = cur
				cur = cur.Left
			} else { // 第二次到达
				LeftRight.Right = nil
				cur = cur.Right
			}
		} else {
			res = append(res, cur.Val)
			cur = cur.Right
		}

	}
	return res
}
func MorrisMid(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	// 定义返回的数组
	res := make([]int, 0)
	// 当前节点
	cur := root
	for cur != nil {
		// 存在左节点
		if cur.Left != nil {
			// 左节点最右的值
			var LeftRight *TreeNode = cur.Left
			// 找到最右节点
			for LeftRight.Right != nil && LeftRight.Right != cur {
				LeftRight = LeftRight.Right
			}
			// 判断是第几次来到当前节点
			if LeftRight.Right == nil { // 第一次到达
				LeftRight.Right = cur
				cur = cur.Left
			} else { // 第二次到达
				res = append(res, cur.Val)
				LeftRight.Right = nil
				cur = cur.Right
			}
		} else {
			res = append(res, cur.Val)
			cur = cur.Right
		}

	}
	return res
}

func MorrisEpil(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	// 定义返回的数组
	res := make([]int, 0)
	// 当前节点
	cur := root
	for cur != nil {
		// 存在左节点
		if cur.Right != nil {
			// 右节点节点最左的值
			var RightLeft *TreeNode = cur.Right
			// 找到最右节点
			for RightLeft.Left != nil && RightLeft.Left != cur {
				RightLeft = RightLeft.Left
			}
			// 判断是第几次来到当前节点
			if RightLeft.Left == nil { // 第一次到达
				res = append(res, cur.Val)
				RightLeft.Left = cur
				cur = cur.Right
			} else { // 第二次到达
				RightLeft.Left = nil
				cur = cur.Left
			}
		} else {
			res = append(res, cur.Val)
			cur = cur.Left
		}

	}
	reverse(res)
	return res
}
func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}
