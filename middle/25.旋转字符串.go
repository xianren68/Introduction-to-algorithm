package middle

func isRevolve(str1, str2 string) bool {
	// 长度不等，直接return
	if len(str1) != len(str2) {
		return false
	}
	// 拼接一个串
	newStr := str1 + str1
	// 通过kmp 算法判断str2是否为newStr的子串
	return kmp(newStr, str2)

}
func kmp(str1, str2 string) bool {
	next := creatNext(str2)
	l := 0
	r := 0
	for r < len(str2) && l < len(str1) {
		if str1[l] == str2[r] {
			l++
			r++
		} else if r > 0 {
			r = next[r]
		} else {
			l++
		}
	}
	if r == len(str2) {
		return true
	}
	return false

}

// 创建前缀数组
func creatNext(str string) []int {
	// 前缀数组
	next := make([]int, len(str))
	next[0] = -1
	next[1] = 0
	// 当前索引
	cur := 2
	// 上一个前缀的下一个值
	nex := 0
	for cur < len(str) {
		if str[cur-1] == str[nex] {
			next[cur] = nex + 1
			cur++
			nex++
		} else if nex > 0 {
			nex = next[nex]
		} else {
			next[cur] = 0
			cur++
		}

	}
	return next
}
