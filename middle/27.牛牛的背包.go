package middle

import (
	"math"
)

func Snack(v []int, sur int) int {
	sum := 0
	for _, val := range v {
		sum += val
	}
	if sum <= sur {
		return int(math.Pow(float64(2), float64(len(v))))
	}
	return Bag(v, 0, sur)
}

// 暴力递归
// v零食，i当前索引，sur，容量剩余
func Bag(v []int, i, sur int) int {
	if i == len(v) {
		return 1
	}
	// 可能会有的两种情况
	//1.不加当前值
	no := Bag(v, i+1, sur)
	//2.判断当前值是否可以加入
	var yes = 0
	if sur-v[i] >= 0 {
		yes = Bag(v, i+1, sur-v[i])
	}
	return no + yes
}
