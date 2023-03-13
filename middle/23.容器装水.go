package middle

func WaterFill(nums []int) int {
	// 至少有三个数据才能装水
	if len(nums) < 3 {
		return 0
	}
	// 左右指针
	left := 1
	right := len(nums) - 1
	// 左右边最大值
	leftMax := nums[0]
	rightMax := nums[len(nums)-1]
	// 收集的水
	res := 0
	for left <= right {
		// 大于等于边界，不可能收集到水，直接指针前进
		if nums[left] >= leftMax {
			leftMax = nums[left]
			left++
			continue
		}
		if nums[right] >= rightMax {
			rightMax = nums[right]
			right--
			continue
		}
		// 哪边小就可以直接确认相加，让指针移动，因为另一边以及有一个更大的值，限制收集的就是较小的边界
		if leftMax < rightMax {
			res += leftMax - nums[left]
			left++
		} else {
			res += rightMax - nums[right]
			right--
		}
	}
	return res

}
