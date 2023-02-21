package basic_sort

func MergeSort(nums []int, left int, right int) []int {
	if left == right {
		return nums
	}
	// 求出中间的位置
	mid := left + (right-left)>>1
	// 递归
	MergeSort(nums, left, mid)
	MergeSort(nums, mid+1, right)
	// 合并两个有序序列
	merge(nums, left, mid, right)
	return nums
}
func merge(nums []int, left, mid, right int) {
	// 定义一个额外的数组
	help := make([]int, len(nums))
	l := left
	// 右边数组开始的位置
	r := mid + 1
	// 合并两个数组
	for i := 0; i <= right-left; i++ {
		// 两边都没有到底
		if l <= mid && r <= right {
			if nums[l] <= nums[r] {
				help[i] = nums[l]
				l++
			} else {
				help[i] = nums[r]
				r++
			}

		} else if l > mid {
			help[i] = nums[r]
			r++
		} else if r > right {
			help[i] = nums[l]
			l++
		}

	}
	j := 0
	// 复制回原数组
	for i := left; i <= right; i++ {
		nums[i] = help[j]
		j++
	}
}
