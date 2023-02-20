package basic_sort

// 选择排序
func SelectSort(nums []int) []int {
	// 判断临界条件
	if nums == nil || len(nums) < 2 {
		return nums
	}
	// 遍历n次，每次找到一个最小值
	for i := 0; i < len(nums)-1; i++ {
		// 最小值的索引
		min := i
		// 从当前值开始遍历
		for j := i; j < len(nums); j++ {
			// 如果此时值比最小值小，则记录此位置为最小
			if nums[j] < nums[min] {
				min = j
			}
		}
		// 交换当前值和最小值的位置
		nums[i], nums[min] = nums[min], nums[i]
	}
	return nums
}
