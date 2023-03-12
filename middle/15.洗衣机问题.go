package middle

import (
	"fmt"
	"math"
)

func findMinMoves(machines []int) int {
	// 判断临界条件
	if len(machines) == 0 {
		return -1
	}
	// 预操作数组，用来记录每个位置的累加和
	sumArr := make([]int, len(machines))
	// 计算累加和以及平均值
	sum := 0
	for i, val := range machines {
		sum += val
		sumArr[i] = sum
	}
	fmt.Println(sumArr)
	// 如果不能除净，则不可能均分
	if remainder := sum % len(machines); remainder != 0 {
		return -1
	}
	// 每个位置上应该的值
	avg := sum / len(machines)
	// 某个位置的次数中的最大值
	max := 0
	for i := 0; i < len(machines); i++ {
		// 左右的需求
		var left, right int
		// 当前值左边需要的衣服
		if i == 0 {
			left = 0
		} else {
			left = sumArr[i-1] - avg*(i)
		}
		// 当前值右边需要的衣服
		right = (sumArr[len(machines)-1] - sumArr[i]) - avg*(len(machines)-1-i)
		// 当前位置需要操作的次数
		var s int
		// 两边都需要衣服，两边相加
		if left < 0 && right < 0 {
			s = int(math.Abs(float64(left)) + math.Abs(float64(right)))
			// 其他情况，绝对值较大的那个
		} else {
			s = int(math.Max(math.Abs(float64(left)), math.Abs(float64(right))))
		}
		// 判断它是否大于最大操作次数
		if s > max {
			max = s
		}
	}
	return max
}

// 优化一波,不用预处理数组，而用两个变量
func findMinMoves2(machines []int) int {
	// 判断临界条件
	if len(machines) == 0 {
		return -1
	}
	// 左边的累加和
	leftSum := 0
	// 右边的累加和
	rightSum := 0
	// 所有数的累加和
	sum := 0
	for _, val := range machines {
		sum += val
	}
	// 如果不能除净，则不可能均分
	if remainder := sum % len(machines); remainder != 0 {
		return -1
	}
	rightSum = sum
	// 每个位置上应该的值
	avg := sum / len(machines)
	// 某个位置的次数中的最大值
	max := 0
	for i := 0; i < len(machines); i++ {
		// 左右的需求
		var left, right int
		// 当前值左边需要的衣服
		left = leftSum - avg*(i)
		// 当前值右边需要的衣服
		right = rightSum - machines[i] - avg*(len(machines)-1-i)
		// 当前位置需要操作的次数
		var s int
		// 两边都需要衣服，两边相加
		if left < 0 && right < 0 {
			s = int(math.Abs(float64(left)) + math.Abs(float64(right)))
			// 其他情况，绝对值较大的那个
		} else {
			s = int(math.Max(math.Abs(float64(left)), math.Abs(float64(right))))
		}
		// 判断它是否大于最大操作次数
		if s > max {
			max = s
		}
		leftSum += machines[i]
		rightSum -= machines[i]
	}
	return max
}
