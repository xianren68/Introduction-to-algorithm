package basic_sort

func InsertSort(nums []int) []int {
	// 判断临界条件
	if nums == nil || len(nums) < 2 {
		return nums
	}
	// 遍历
	for i := 0; i < len(nums)-1; i++ {
		// 比较当前值与前面的有序序列
		for j := i + 1; j > 0; j-- {
			// 如果当前值较小则与前面的值换位
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
	return nums
}
