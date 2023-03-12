package middle

func rotate(matrix [][]int) {
	if len(matrix[0]) == 1 {
		return
	}
	leftX := 0
	leftY := 0
	// 右下角顶点坐标
	rightX := len(matrix[0]) - 1
	rightY := len(matrix) - 1
	for leftX <= rightX && leftY <= rightY {
		// 每一层分组交换
		// 组数
		s := rightX - leftX
		// 用于交换的空间
		var tmp int
		// 换位
		for i := 0; i < s; i++ {

			tmp = matrix[rightY-i][leftX]
			matrix[rightY-i][leftX] = matrix[rightY][rightX-i]
			matrix[rightY][rightX-i] = matrix[leftY+i][rightX]
			matrix[leftY+i][rightX] = matrix[leftY][leftX+i]
			matrix[leftY][leftX+i] = tmp
		}
		// 移动顶点
		leftX++
		leftY++
		rightY--
		rightX--
	}

}
