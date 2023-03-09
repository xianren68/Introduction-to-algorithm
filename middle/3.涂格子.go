package middle

import (
	"fmt"
	"math"
)

func DrawGrid(s string) int {
	// 预处理数组，判断每个位置左边R的个数
	RArr := make([]int, len(s))
	GArr := make([]int, len(s))
	r := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			r += 1
		}
		RArr[i] = r
	}
	g := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'G' {
			g += 1
		}
		GArr[i] = g
	}
	fmt.Println(RArr, GArr)
	min := math.MaxInt
	for i := 0; i < len(s); i++ {
		s := RArr[i] + GArr[i] - 1
		if s < min {
			min = s
		}
	}
	return min
}
