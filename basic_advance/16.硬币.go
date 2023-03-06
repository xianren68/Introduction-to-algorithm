package basic_advance

import (
	"math"
)

func CoinViolence(nums []int, aim int) int {
	return CProcess(nums, 0, aim)
}

// index,当前硬币所在的位置
// aim，到目标还剩多少
func CProcess(nums []int, index int, aim int) int {
	if aim < 0 {
		return -1
	}
	if aim == 0 {
		return 0
	}
	// 到达末位，返回
	if index == len(nums) {
		return -1
	}
	// 第一种可能，自己本身不加入
	p1 := CProcess(nums, index+1, aim)
	// 第二种可能，自己本身加入
	p2 := CProcess(nums, index+1, aim-nums[index])
	// 没有成功的结果
	if p1 == -1 && p2 == -1 {
		return -1
	}
	if p1 == -1 {
		return p2 + 1
	}
	if p2 == -1 {
		return p1
	}
	// 成功的结果取最小的
	return int(math.Min(float64(p1), float64(p2+1)))

}
