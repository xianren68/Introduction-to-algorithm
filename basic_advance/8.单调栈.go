package basic_advance

func Monotonous(nums []int) [][2]int {
	if len(nums) == 0 {
		return nil
	}
	// 定义一个数组用作单调栈
	// 为了兼容有相同值情况直接用数组来存储
	stack := make([][]int, 0)
	// 返回的数组,第一位为后面的大值，第二位为前面的大值，-1为不存在
	res := make([][2]int, len(nums))
	for i, val := range nums {
		// 将小于当前值的全部出栈
		for len(stack) > 0 && nums[stack[len(stack)-1][0]] < val {
			arr := stack[len(stack)-1]
			for _, valu := range arr {
				//后面的值
				if len(stack) > 1 {
					// 栈中底下的位置
					res[valu][0] = nums[stack[len(stack)-2][0]]
				} else {
					res[valu][0] = -1
				}
				// 后面的值等于后面要替换的值
				res[valu][1] = val

			}
			// 栈长度-1
			stack = stack[:len(stack)-1]
		}
		// 如果相等，则加入数组
		if len(stack) > 0 && val == nums[stack[len(stack)-1][0]] {
			stack[len(stack)-1] = append(stack[len(stack)-1], i)
			continue
		}
		// 不相等，直接加入索引
		stack = append(stack, []int{i})
	}
	// 弹出栈中剩下的值
	for len(stack) > 0 {
		arr := stack[len(stack)-1]
		for _, val := range arr {
			// 没有这个值
			res[val][1] = -1
			if len(stack) > 1 {
				res[val][0] = nums[stack[len(stack)-2][0]]
			} else {
				res[val][0] = -1
			}
		}
		stack = stack[:len(stack)-1]
	}
	return res
}
