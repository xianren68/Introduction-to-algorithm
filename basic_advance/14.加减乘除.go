package basic_advance

// 加
func Add(a, b int) int {
	for b != 0 {
		// 无进位相加
		sum := a ^ b
		// 进位信息
		b = (a & b) << 1
		a = sum
	}
	return a
}

// 取相反数
func negNum(num int) int {
	return Add(^num, 1)
}

// 减
func Sub(a, b int) int {
	return Add(a, negNum(b))
}

// 乘
func Ride(a, b int) int {
	res := 0
	for b != 0 {
		// 最后一位为1,相加
		if (b & 1) != 0 {
			res = Add(res, a)
		}
		a = a << 1
		b = b >> 1

	}
	return res
}
