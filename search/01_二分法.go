package search

import "fmt"

// 二分法，升序列表查值
func Dichotomous(nums []int, s int) {
	// 定义左右两个指针
	x := 0
	y := len(nums) - 1
	for x <= y {
		// 中间位置
		mid := (y-x)/2 + x
		if nums[mid] == s {
			fmt.Println(mid)
			return
		} else if nums[mid] > s {
			y = mid - 1
		} else {
			x = mid + 1
		}
	}
	fmt.Println(x)
}

// // 局部最小值
// func LocalMin(nums []int) {
// 	// 首先判断两端
// }
