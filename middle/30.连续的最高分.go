package middle

import "math"

func HightScore(nums []int) int {
	// 记录最大值
	max := math.MinInt
	// 记录累加
	cur := 0
	// 遍历数组
	for i := 0; i < len(nums); i++ {
		cur += nums[i]
		if cur > max {
			max = cur
		}
		if cur <= 0 {
			cur = 0
		}

	}
	return max
}
