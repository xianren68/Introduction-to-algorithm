package basic_advance

import (
	"math"
)

// Manacher
func Manacher(str string) int {
	if len(str) == 0 {
		return 0
	} else if len(str) == 1 {
		return 1
	}
	// 将字符串转为字节数组,并插入特殊字符#
	strArr := make([]byte, len(str)<<1)
	j := 0
	for i := 0; i < len(strArr); i += 2 {
		strArr[i] = '#'
		if i < len(strArr) {
			strArr[i+1] = str[j]
			j++
		}
	}
	strArr = append(strArr, '#')
	// 定义最大半径数组
	MaxRadius := make([]int, len(strArr))
	// 最大边界对应的中心
	C := -1
	// 定义最大右边界
	R := -1
	// 定义最大值
	var max = math.MinInt
	// 遍历
	for i := 0; i < len(strArr); i++ {
		// i至少的回文区域
		if i >= R {
			// i小于等于R,只有自己是确定的，其它的得扩展
			MaxRadius[i] = 1
		} else {
			// R>i，最少长度为i-R的距离或者i'的回文长度
			MaxRadius[i] = int(math.Min(float64(R-i), float64(MaxRadius[C<<1-i])))
		}
		// 向两边扩展，直到值不相同或到达数组边界
		for i+MaxRadius[i] < len(strArr) && i-MaxRadius[i] > -1 {
			if strArr[i+MaxRadius[i]] == strArr[i-MaxRadius[i]] {
				MaxRadius[i]++ //增加长度
			} else {
				break
			}
		}
		// 判断当前值的最大边界
		// 如果超越了旧的则直接替换
		if i+MaxRadius[i] > R {
			R = i + MaxRadius[i]
			C = i
		}
		max = int(math.Max(float64(MaxRadius[i]), float64(max)))
	}
	// 返回子序列最大长度，这里加了特殊字符所以-1的半径就是原本的长度
	return max - 1

}
