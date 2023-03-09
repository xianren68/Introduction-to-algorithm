package middle

import (
	"algorithm/utils"
	"math"
	"time"
)

// 返回a - b之间的随机数函数
func Rand(a, b int) func() int {
	return func() int {
		return utils.RandomInt(b) + a
	}
}

// 1
func RandomSeven() int {
	// 7是3位，随机三个二进制拼起来
	r := 0
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		// 1-5的随机数方法
		x := OneOrTwo(1, 5)
		r += x * int(math.Pow(float64(2), float64(i)))
	}
	return r
}
func OneOrTwo(a, b int) int {
	// 已有的a-b的随机方法,将其转化为返回0/1
	f := Rand(a, b)
	if f() == 1 {
		return OneOrTwo(a, b)
	}
	if f() < 3 {
		return 0
	}
	return 1
}
