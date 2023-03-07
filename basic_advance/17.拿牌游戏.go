package basic_advance

import (
	"math"
)

// 纸牌游戏暴力递归
func PlateViolence(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	// 先手或后手中最大的一个
	p := pre(nums, 0, len(nums)-1)
	c := cur(nums, 0, len(nums)-1)
	return int(math.Max(float64(p), float64(c)))
}

// l和r代表左右的值
// 先手情况
func pre(nums []int, l, r int) int {
	// 取到最后一张
	if r == l {
		return nums[l]
	}
	// 取一张对自己最大的值
	p := nums[l] + cur(nums, l+1, r)
	c := nums[r] + cur(nums, l, r-1)
	return int(math.Max(float64(p), float64(c)))
}

// 后手情况
func cur(nums []int, l, r int) int {
	// 后手，没办法选了
	if l == r {
		return 0
	}
	// 先手的人拿了左边的值
	p := pre(nums, l+1, r)
	// 先手的人拿了右边的值
	c := pre(nums, l, r-1)
	return int(math.Min(float64(p), float64(c)))
}

// 严格表模式
func StrictDp(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	// 创建两张二维表
	p := make([][]int, len(nums))
	c := make([][]int, len(nums))
	for i, _ := range nums {
		p[i] = make([]int, len(nums))
		c[i] = make([]int, len(nums))
	}
	// 根据递归条件得，先手表在l和r相等时是nums[l]的值
	// 后手表在两者相等时为0
	for i, _ := range p {
		p[i][i] = nums[i]
	}
	// 根据观察和依赖分析，可以发现两张表是相互依赖的
	// 从第一行第二列开始沿着斜率遍历，因为l<r所以另外半张表没用
	row := 0
	col := 1
	// 斜着遍历
	for col < len(nums) {
		i := row
		j := col
		//
		for i < len(nums) && j < len(nums) {
			p[i][j] = int(math.Max(float64(nums[i]+c[i+1][j]), float64(nums[j]+c[i][j-1])))
			c[i][j] = int(math.Min(float64(p[i+1][j]), float64(p[i][j-1])))
			i++
			j++
		}
		col++
	}
	return int(math.Max(float64(p[0][len(nums)-1]), float64(c[0][len(nums)-1])))
}
