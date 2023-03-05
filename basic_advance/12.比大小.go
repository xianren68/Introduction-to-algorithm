package basic_advance

func Compare(a int, b int) int {
	c := a - b
	// 值为正数或负数
	scA := isPositive(c)
	// 取得相反的数
	scB := flip(scA)
	// scA与scB是互斥的，一个为1的时候另一个为0
	return a*scA + b*scB

}

// 将一个0/1转化为1/0
func flip(i int) int {
	return i ^ 1
}

// 判断一个32位整数是否为正数，返回1为正数，0为负数
func isPositive(c int) int {
	// 先将数右移31位令其最左边的值到最右边，其余位置全为0
	return flip(c >> 31 & 1)
}
