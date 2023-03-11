package middle

func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	// 一个数组，用来装以每个位置为末位的最长子串
	arr := make([]int, len(s))
	// 记录最大值
	max := 0
	// 第一个数必然为0
	arr[0] = 0
	for i := 1; i < len(s); i++ {
		// 以它为结尾，结果为0
		if s[i] == '(' {
			arr[i] = 0
		} else {
			// 前面一个为'(',总长为2
			if arr[i-1] == 0 && s[i-1] == '(' {
				// 加上前面可能的有效子串
				if i-2 >= 0 {
					arr[i] = 2 + arr[i-2]
				} else {
					arr[i] = 2
				}
				// 前面有有效串，判断这个有效串之前的一个元素是否为'('
			} else if i-arr[i-1] >= 1 && s[i-arr[i-1]-1] == '(' {
				// 如果是并且前面还有数，则加上前面的数
				if i-arr[i-1]-1 > 1 {
					arr[i] = arr[i-1] + 2 + arr[i-arr[i-1]-2]
				} else {
					arr[i] = arr[i-1] + 2
				}
			} else {
				arr[i] = 0
			}

		}
		if max < arr[i] {
			max = arr[i]
		}

	}
	return max
}
