package middle

func isExist(nums [][]int, num int) bool {
	// 从第一行的末位开始遍历
	i := 0
	j := len(nums[0])
	for (i >= 0 && i < len(nums)) && (j >= 0 && j < len(nums[0])) {
		if nums[i][j] == num {
			return true
		} else if nums[i][j] > num {
			j--
		} else if nums[i][j] < num {
			i++
		}
	}
	return false
}
