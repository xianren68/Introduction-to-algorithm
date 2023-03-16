package middle

// 动态规划
func lengthOfLIS(nums []int) int {
	// 用一个数组，记录以每个数为末位的连续递增的长度
	arr := make([]int, len(nums))
	// 将第一个值加入
	arr[0] = 1
	// 记录最大长度
	max := 1

	for i := 1; i < len(nums); i++ {
		m := 1
		for j := i - 1; j >= 0; j-- {
			// 满足递增
			if nums[i] > nums[j] {
				// 判断长度是否大于m
				if arr[j]+1 > m {
					m = arr[j] + 1
				}
			}
		}
		arr[i] = m
		if m > max {
			max = m
		}
	}
	return max
}
