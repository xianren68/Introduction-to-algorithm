package middle

import "math"

func MaxValue(nums []int) int {
	// 找到整个数组的最大值
	max := math.MinInt
	for _, val := range nums {
		if val > max {
			max = val
		}
	}
	// 找到两边较小的一个
	var min = nums[0]
	if nums[0] > nums[len(nums)-1] {
		min = nums[len(nums)-1]
	}
	// 相减的绝对值
	return int(math.Abs(float64(max - min)))

}
