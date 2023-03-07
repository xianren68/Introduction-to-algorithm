package basic_advance

import (
	"math"
	"strconv"
)

// 暴力递归
func BobViolence(n, m, i, j, k int) string {
	// 鲍勃可以尝试的走法总次数
	total := int(math.Pow(4, float64(k)))
	live := bobProcess(n, m, i, j, k)
	return strconv.Itoa(live) + "/" + strconv.Itoa(total)
}

// 递归方法
func bobProcess(n, m, i, j, k int) int {
	// 如果已经在外面，没救了
	if i >= n || j >= m || i < 0 || j < 0 {
		return 0
	}
	// 在里面，并且不用走了
	if k == 0 {
		return 1
	}
	// 在里面还需要走
	// 有上下左右四种走法
	return bobProcess(n, m, i-1, j, k-1) +
		bobProcess(n, m, i+1, j, k-1) +
		bobProcess(n, m, i, j+1, k-1) +
		bobProcess(n, m, i, j-1, k-1)
}

// 严格表
func StrictBob(n, m, i, j, k int) int {
	// 创建三维数组，底面为i和j,高为剩余步数k
	Dp := make([][][]int, k+1)
	for i, _ := range Dp {
		Dp[i] = newSlice(n, m)
	}
	for i, _ := range Dp {
		for j, _ := range Dp[i] {
			// 没有步数还在范围内就是一种存活的可能
			Dp[i][j][0] = 1
		}
	}
	for h := 1; h <= k; h++ {
		for row := 0; row < n; row++ {
			for col := 0; col < m; col++ {
				Dp[row][col][h] = getValue(n, m, row-1, col, h-1, Dp) +
					getValue(n, m, row+1, col, h-1, Dp) +
					getValue(n, m, row, col-1, h-1, Dp) +
					getValue(n, m, row, col+1, h-1, Dp)
			}
		}
	}
	return Dp[i][j][k]

}

// 二维数组
func newSlice(n, m int) [][]int {
	res := make([][]int, n+1)
	for i, _ := range res {
		res[i] = make([]int, m+1)
	}
	return res
}

// 获取数据避免越界
func getValue(n, m, i, j, k int, Dp [][][]int) int {
	if i < 0 || j < 0 || i >= n || j >= m {
		return 0
	}
	return Dp[i][j][k]
}
