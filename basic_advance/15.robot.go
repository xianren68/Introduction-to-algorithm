package basic_advance

// 1.暴力递归
func RobotViolence(n, s, e, k int) int {
	// n 总长度
	// s 当前位置
	// e 要到达的地方
	// k 还剩的步数
	return Rprocess1(n, e, s, k)
}

// 递归函数
func Rprocess1(n, e, s, k int) int {
	if k == 0 {
		if s == e {
			return 1
		}
		return 0
	}
	// 来到了开头位置，只能前进
	if s == 1 {
		return Rprocess1(n, e, s+1, k-1)
	}
	// 来到了末位，只能后退
	if s == n {
		return Rprocess1(n, e, s-1, k-1)
	}
	return Rprocess1(n, e, s+1, k-1) + Rprocess1(n, e, s-1, k-1)
}

// 2.记忆搜索
func RobotRem(n, e, s, k int) int {
	// 创建二维数组
	remArr := make([][]int, n+1)
	for i, _ := range remArr {
		remArr[i] = make([]int, k+1)
	}
	// 将数组用-1填充
	for i, _ := range remArr {
		for j, _ := range remArr[i] {
			remArr[i][j] = -1
		}
	}
	return Rprocess2(n, e, s, k, remArr)
}

func Rprocess2(n, e, s, k int, nums [][]int) int {
	if nums[s][k] != -1 {
		return nums[s][k]
	}
	if k == 0 {
		if s == e {
			nums[s][k] = 1
		} else {
			nums[s][k] = 0
		}
		return nums[s][k]
	}
	if s == 1 {
		nums[s][k] = Rprocess2(n, e, s+1, k-1, nums)
	} else if s == n {
		nums[s][k] = Rprocess2(n, e, s-1, k-1, nums)
	} else {
		nums[s][k] = Rprocess2(n, e, s+1, k-1, nums) + Rprocess2(n, e, s-1, k-1, nums)
	}
	return nums[s][k]
}

// 3. 严格表
func StrictRobot(n, e, s, k int) int {
	// 创建二维表
	// 根据前面的递归决定依赖关系
	// 创建二维数组
	remArr := make([][]int, k+1)
	for i, _ := range remArr {
		remArr[i] = make([]int, n+1)
	}
	// 步数为0时只有e位置为1
	remArr[0][e] = 1
	// 位置无法到达0，位置为1时它的值来自它的位置向前，步数减1的可能，正好再数组右上角
	// 位置到n时，它对应的值为左上角
	// 位置在其他位置时，对应的可能性值为左上角+右上角
	// 遍历赋值
	for i := 1; i <= k; i++ {
		for j := 1; j <= n; j++ {

			if j == 1 {
				remArr[i][j] = remArr[i-1][j+1]
			} else if j == n {
				remArr[i][j] = remArr[i-1][j-1]
			} else {
				remArr[i][j] = remArr[i-1][j+1] + remArr[i-1][j-1]
			}
			if i == k && j == s {
				return remArr[i][j]
			}
		}
	}
	return -1

}
