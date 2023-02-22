package basic_sort

func BubbleSort(nums []int) []int {
	// 判断临界条件
	if nums == nil || len(nums) < 2 {
		return nums
	}
	// 比较的轮数
	for i := 0; i < len(nums)-1; i++ {
		// 遍历，比较相邻两个数的大小
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				// 换位置
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
