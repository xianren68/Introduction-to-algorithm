package basic_advance

// 是不是2的幂
func IsTwoPower(num int) bool {
	// 2的幂所有位上只有一个1
	// -1后可以让那个1那个位变成0，它后面的值全为1
	// 它们与之后等于0
	return (num & (num - 1)) == 0
}

// 是不是4的幂
func IsFourPower(num int) bool {
	// 判断是不是2的幂
	if !IsTwoPower(num) {
		return false
	}
	// 4的幂那个特殊的1只能出现在奇数位上
	// 找一个形如010101010101的值与当前值与即可
	return (num & 0x55555555) != 0
}
