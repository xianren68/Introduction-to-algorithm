package basic_sort

import (
	"algorithm/utils"
	"fmt"
)

// 快速排序
func QuickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	// 生成一个随机数作为num的位置
	index := utils.RandomInt(right-left) + left
	// 交换当前值和最后一个值
	nums[index], nums[right] = nums[right], nums[index]
	// 将值排成三个区域
	l, r := partition(nums, left, right)
	fmt.Println(left, right, index)
	fmt.Println(l, r)
	// 递归
	QuickSort(nums, left, l-1)
	QuickSort(nums, r+1, right)
	fmt.Println(nums)

}

// 荷兰国旗问题，用于将数组分为三部分
func partition(nums []int, left, right int) (int, int) {
	// 双指针
	// l代表<区，r代表>区，中间为等于区
	l := left
	r := right - 1
	// 最后一个值
	last := nums[right]
	i := left
	for i <= r {
		// 小于给定值,交换到<区的后一个位置，扩展<区
		if nums[i] < last {
			nums[i], nums[l] = nums[l], nums[i]
			i++
			l++
			// 大于给定值，不发生变化,将当前值交换到>区的前一个位置，拓展>区
		} else if nums[i] > last {
			nums[r], nums[i] = nums[i], nums[r]
			r--
		} else {
			i++
		}

	}
	// 将最后一个值换到中间区域
	nums[r+1], nums[right] = nums[right], nums[r+1]
	r++
	return l, r
}
