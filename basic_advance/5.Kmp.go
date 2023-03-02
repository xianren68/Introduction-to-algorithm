package basic_advance

func Kmp(str1, str2 string) int {
	if len(str2) > len(str1) {
		return -1
	}
	next := creatNext(str2)
	l := 0
	n := 0
	for n < len(str2) && l < len(str1) {
		// 相等则两个都往后走
		if str1[l] == str2[n] {
			l++
			n++
			// 不相同，并且n>0，还可以往前跳
		} else if n > 0 {
			n = next[n]
			// n == 0了，直接开始下一个值的对比
		} else {
			l++

		}
	}
	// 判断是谁到达末位
	if n == len(str2) {
		return l - n
	}
	return -1

}
func creatNext(str1 string) []int {
	if len(str1) == 1 {
		return []int{-1}
	}
	//定义返回的数组
	next := make([]int, len(str1))
	next[0] = -1
	next[1] = 0
	// 当前索引
	i := 2
	// 前缀的下一个位置
	cur := 0
	for i < len(str1) {
		// 前一个值的最长前缀的下一个值与前一个值相等，直接下一个值
		if str1[cur] == str1[i-1] {
			next[i] = cur + 1
			i++
			cur++
			// 还大于0还可以继续往前跳
		} else if cur > 0 {
			cur = next[cur]
		} else {
			next[i] = 0
			i++
		}
	}
	return next
}
