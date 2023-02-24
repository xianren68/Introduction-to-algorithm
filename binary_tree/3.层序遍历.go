package binary_tree

func levelOrder(root *TreeNode) [][]int {
	// 判断是否为空
	if root == nil {
		return make([][]int, 0)
	}
	// 返回的数组
	res := make([][]int, 0)
	// 用一个切片来模拟队列
	quene := make([]*TreeNode, 0)
	// 头节点入队
	quene = append(quene, root)
	for len(quene) > 0 {
		// 获取当前队列的长度(一层的节点个数)
		size := len(quene)
		// 数组用来存储值
		arr := make([]int, 0)
		for i := 0; i < size; i++ {
			// 出队
			node := quene[0]
			quene = quene[1:]
			// 将数加入数组
			arr = append(arr, node.Val)
			// 判断左右节点是否为空,不为空则入队
			if node.Left != nil {
				quene = append(quene, node.Left)
			}
			if node.Right != nil {
				quene = append(quene, node.Right)
			}
		}
		// 将每一层的数组加入二维数组
		res = append(res, arr)
	}
	return res
}
