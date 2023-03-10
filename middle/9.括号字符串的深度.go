package middle

func Depth(s string) int {
	// 记录最大深度
	max := 0
	// 记录多余的'("
	cur := 0
	for _, val := range s {
		if val == '(' {
			if cur == max {
				max++
			}
			cur++
		} else {
			cur--
		}
	}
	return max
}
