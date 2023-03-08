package basic_advance

func CoinViolence2(nums []int, aim int) int {
	return coinProcess2(nums, 0, aim)
}

// 暴力递归
func coinProcess2(nums []int, index int, aim int) int {
	if aim < 0 {
		return 0
	}
	if index == len(nums) {
		if aim == 0 {
			return 1
		}
		return 0
	}
	ways := 0
	for i := 0; nums[index]*i <= aim; i++ {
		ways += coinProcess2(nums, index+1, aim-i*nums[index])
	}
	return ways
}
