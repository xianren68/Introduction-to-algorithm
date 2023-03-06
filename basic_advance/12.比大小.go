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

// 上面写的方法还有一个问题就是如果符号不同的情况下，相减可能会出现溢出
func Compare2(a, b int) int {
	c := a - b
	// 得到a和b的正负
	sa := isPositive(a)
	sb := isPositive(b)
	sc := isPositive(c)
	// ab 异号为1
	difSab := sa ^ sb
	// ab同号为1
	isSame := flip(difSab)
	// 如果两者异号，则返回a的正负号，正则为1，如果两者同号则返回两者相减后的值的正负号
	returnA := difSab*sa + isSame*sc
	returnB := flip(returnA)
	return a*returnA + b*returnB

}
