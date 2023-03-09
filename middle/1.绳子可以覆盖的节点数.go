package middle

func Cord(nums []int, L int) int {
	if L == 0 {
		return 0
	}
	// 窗口左右边界
	l := 0
	r := 0
	// 最大个数
	max := 0
	// 窗口内的个数
	s := 0
	// 开始遍历后面的位置
	for l <= r && r < len(nums) {
		// 边界相等，长度为1
		if l == r && l < len(nums)-1 {
			// 如果数据差距过大，r跟随l一同跨过去，防止r卡住
			if nums[l]+L < nums[l+1] {
				l++
				r = l
			}
			s = 1
			// 左边界扩展，个数-1
		} else {
			l++
			s--
		}
		// 右边界扩展,找到绳子最后能碰到的位置停下
		for r < len(nums)-1 {
			if nums[r+1]-nums[l] <= L {
				r++
				s++
			} else {
				break
			}
		}
		if s > max {
			max = s
		}
	}
	return max
}
