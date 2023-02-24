package binary_tree

// 1.先序
func PreTraversal(root *TreeNode) []int {
	// 准备返回的数组
	res := make([]int, 0)
	if root == nil {
		return res
	}
	// 准备一个栈
	stack := make([]*TreeNode, 0)
	// 头节点入栈
	stack = append(stack, root)
	// 依次出入栈
	for len(stack) != 0 {
		// 出栈，并加入数组
		val := stack[len(stack)-1]
		// 栈长度-1
		stack = stack[:len(stack)-1]
		res = append(res, val.Val)
		if val.Right != nil {
			stack = append(stack, val.Right)
		}
		if val.Left != nil {
			// 入栈
			stack = append(stack, val.Left)
		}

	}
	return res
}
func InTraversal(root *TreeNode) []int {
	// 准备返回的数组
	res := make([]int, 0)
	if root == nil {
		return res
	}
	// 准备两个个栈
	stack1 := make([]*TreeNode, 0)
	stack2 := make([]*TreeNode, 0)
	// 头节点入栈
	stack1 = append(stack1, root)
	// 依次出入第一个栈
	for len(stack1) != 0 {
		// 出栈，并加入数组
		val := stack1[len(stack1)-1]
		// 栈长度-1
		stack1 = stack1[:len(stack1)-1]
		// 值压入第二个栈
		stack2 = append(stack2, val)
		if val.Left != nil {
			// 入栈
			stack1 = append(stack1, val.Left)
		}
		if val.Right != nil {
			stack1 = append(stack1, val.Right)
		}

	}
	// 依次出栈
	for len(stack2) != 0 {
		// 出栈，并加入数组
		val := stack2[len(stack2)-1]
		// 栈长度-1
		stack2 = stack2[:len(stack2)-1]
		// 将数存入数组
		res = append(res, val.Val)
	}
	return res
}
func PostTraversal(root *TreeNode) []int {
	// 准备返回的数组
	res := make([]int, 0)
	if root == nil {
		return res
	}
	// 准备一个栈
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for root != nil {
		// 走到左边界
		for root.Left != nil {
			// 有入栈
			stack = append(stack, root.Left)
			// 往左走
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		for root.Right == nil {
			if len(stack) == 0 {
				return res
			}
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, root.Val)
		}
		root = root.Right
		stack = append(stack, root)
	}
	return res
}
