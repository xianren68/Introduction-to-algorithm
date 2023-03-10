package middle

func LeastBracket(s string) int {
	// 括号个数
	sum := 0
	// 当前所在位置括号的排列
	cur := 0
	for _, val := range s {
		if val == '(' {
			cur++
		} else {
			cur--
		}
		if cur < 0 {
			sum++
			cur = 0
		}
	}
	return cur + sum
}
