package middle

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 统计左子树的深度
	node := root
	Lheight := 0
	for node != nil {
		node = node.LeftNode
		Lheight++
	}
	Lheight--
	// 判断右子树的深度
	Rheight := 0
	node = root.LeftNode
	for node != nil {
		node = node.LeftNode
		Rheight++
	}
	// 左右节点个数
	l := 0
	r := 0
	// 左子树是满的
	if Lheight == Rheight {
		l = 1<<Lheight - 1
		r = countNodes(root.RightNode)
		// 右子树是满的
	} else {
		r = 1<<Rheight - 1
		l = countNodes(root.LeftNode)
	}
	// 返回总个数
	return l + r + 1

}
