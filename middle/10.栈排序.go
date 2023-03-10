package middle

import (
	"fmt"
	"math"
)

func StackSort(nums []int) {
	// 申请一个新栈
	newStack := make([]int, 0)
	newStack = append(newStack, math.MaxInt)
	// 遍历前一个栈
	for len(nums) != 0 {
		num := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		for newStack[len(newStack)-1] <= num {
			// 将大的数重新挪回前面的栈
			nums = append(nums, newStack[len(newStack)-1])
			newStack = newStack[:len(newStack)-1]
		}
		newStack = append(newStack, num)
		// 第一个栈数量-1
	}
	// 将数据重新换回来
	for len(newStack) > 1 {
		nums = append(nums, newStack[len(newStack)-1])
		newStack = newStack[:len(newStack)-1]
	}
	fmt.Println(nums)
}
