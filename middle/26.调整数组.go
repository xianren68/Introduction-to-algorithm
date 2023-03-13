package middle

func Adjust(nums []int) bool {
	// 奇数,2因子，4因子的个数
	var k, l, t int
	for _, val := range nums {
		if val%2 != 0 {
			k++
		} else if val%4 == 0 {
			t++
		} else {
			l++
		}
	}
	if k == 0 && (l%2 == 0 || l == t) {
		return true
	}
	if t == 0 && l%2 == 0 && k == 0 {
		return true
	}
	if l == 0 && t == k {
		return true
	}
	if t == k && t != 0 {
		return true
	}
	return false
}
