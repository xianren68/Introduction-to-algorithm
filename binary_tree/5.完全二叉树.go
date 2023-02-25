package binary_tree

func Complete(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	// 一个队列，用来中序遍历
	quene := make([]*TreeNode, 0)
	// 用来记录前方是否已有节点只存在左节点
	flag := false
	// 头节点入队
	quene = append(quene, root)
	for len(quene) != 0 {
		// 出队节点
		root = quene[0]
		quene = quene[1:]
		// 判断是否满足条件
		if root.Left == nil && root.Right != nil {
			return false
		}
		if !flag {
			if root.Left != nil && root.Right == nil {
				// 节点入栈
				// 只有左节点，标记变为true
				quene = append(quene, root.Left)
				flag = true
				// 左右双全，入栈
			} else if root.Left != nil && root.Right != nil {
				quene = append(quene, root.Left)
				quene = append(quene, root.Right)
			} else {
				flag = true
			}
		} else {
			// 不满足叶子节点
			if root.Left != nil && root.Right != nil {
				return false
			}
		}

	}
	return true
}
