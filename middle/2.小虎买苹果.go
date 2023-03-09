package middle

import "fmt"

// 普通方法
func BuyApple(n int) int {
	// 奇数，不可能满足条件
	if n%2 != 0 || n < 6 {
		return -1
	}
	// 先尽量用8的袋子装
	num := n / 8
	// 剩余的数量
	surplus := n % 8
	for surplus < 24 && surplus != 0 {
		if surplus >= 6 {
			l := surplus / 6
			if surplus%6 == 0 {
				return l + num
			}
		}
		if num > 0 {
			num--
			surplus += 8
		} else {
			return -1
		}

	}
	if surplus == 0 {
		return num
	} else {
		return -1
	}
}

// 打表法
func Table(n int) int {
	// 奇数返回-1
	if n%2 != 0 {
		return -1
	}
	//18以下
	if n < 18 {
		if n == 6 || n == 8 {
			return 1
		} else if n == 12 || n == 16 || n == 14 {
			return 2
		} else {
			return -1
		}
	}
	return (n-18)/8 + 3
}
func Test() {
	for i := 0; i < 100; i += 2 {
		fmt.Println(i, Table(i))
	}
}
