package middle

import "strconv"

func translateNum(num int) int {
	// 先将给定的数字转化为一个字符串
	str := strconv.Itoa(num)
	return process(str, 0)

}
func process(str string, i int) int {
	if i == len(str) {
		return 1
	}
	// 每个值有两种选择，自己成一位或与后面的相加
	one := process(str, i+1)
	// 后面还有
	if i+1 < len(str) && str[i] != '0' {
		// 判断相加是否会大于25
		nu, _ := strconv.Atoi(str[i : i+2])
		if nu <= 25 {
			two := process(str, i+2)
			return one + two
		}
		return one
	} else {
		return one
	}
}
