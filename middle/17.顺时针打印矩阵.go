package middle

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	// 左上角顶点坐标
	leftX := 0
	leftY := 0
	// 右下角顶点坐标
	rightX := len(matrix[0]) - 1
	rightY := len(matrix) - 1
	// 定义返回的数组
	res := make([]int, 0)

	for leftX <= rightX && leftY <= rightY {
		// 遍历打印四个边
		// 上
		for i := leftX; i < rightX; i++ {
			res = append(res, matrix[leftY][i])
		}
		// 右
		for i := leftY; i <= rightY; i++ {
			res = append(res, matrix[i][rightX])
		}
		// 不在同一行
		if rightY != leftY {
			// 下
			for i := rightX - 1; i > leftX; i-- {
				res = append(res, matrix[rightY][i])
			}
		}

		//不在同一列
		if rightX != leftX {
			for i := rightY; i > leftY; i-- {
				res = append(res, matrix[i][leftX])
			}
		}
		// 开始进入下一层
		leftX++
		leftY++
		rightX--
		rightY--
	}
	return res
}
