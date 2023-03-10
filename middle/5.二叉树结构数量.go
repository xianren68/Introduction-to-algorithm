package middle

// 定义二叉树结构
type TreeNode struct {
	Val       int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

// 暴力递归
func TreeQuantity(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	if n == 2 {
		return 2
	}
	// 总排列
	var x = 0
	for i := 0; i < n-1; i++ {
		// 左右可能的排列
		// 相乘即为一次遍历的排列
		left := TreeQuantity(i)
		right := TreeQuantity(n - i - 1)
		x += left * right
	}
	return x
}

// 动态规划
func TreeQuantity2(n int) int {
	// 只有一个可变参数建立一个一维表
	table := make([]int, n+1)
	// 根据递归，先将0值加入
	table[0] = 1
	table[1] = 1
	table[2] = 2
	if n < 3 {
		return table[n]
	}
	// 遍历表
	for i := 3; i < n+1; i++ {
		for j := 0; j < i-1; j++ {
			table[i] += table[j] * table[i-j-1]
		}
	}
	return table[n]
}
