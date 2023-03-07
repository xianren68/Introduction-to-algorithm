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

// 严格表
func StrictCoin(nums []int, aim int) int {
	// 目标小于0,没有结果
	if aim < 0 {
		return -1
	}
	// 创建一个二维表
	table := make([][]int, len(nums)+1)
	for i, _ := range table {
		table[i] = make([]int, aim+1)
		table[i][0] = 0
	}
	for i := 1; i <= aim; i++ {
		table[len(nums)][i] = -1
	}
	// 改写递归
	for i := len(table) - 2; i >= 0; i-- {
		for a := 1; a <= aim; a++ {
			// 自己不参加
			p1 := table[i+1][a]
			// 自己参加
			p2 := -1
			if a-nums[i] >= 0 {
				p2 = table[i+1][a-nums[i]]
			}
			if p1 == -1 && p2 == -1 {
				table[i][a] = -1
			} else if p1 == -1 {
				table[i][a] = p2 + 1
			} else if p2 == -1 {
				table[i][a] = p1
			} else {
				table[i][a] = int(math.Min(float64(p1), float64(p2+1)))
			}

		}
	}
	return table[0][aim]

}
