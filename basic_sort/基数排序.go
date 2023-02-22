package basic_sort

import (
	"math"
)

// 桶排序
func BucketSort(nums []int, base int) {
	if nums == nil || len(nums) < 2 {
		return
	}

	// 准备一个辅助数组
	help := make([]int, len(nums))
	// 获得最大位数
	flag := maxBit(nums)
	// 遍历排序
	// 有多少位就遍历多少次
	for i := 0; i < flag; i++ {
		// 准备一个进制长度的数组
		buckets := make([]int, base)
		// 遍历整个数组，对应位置值++
		for j := 0; j < len(nums); j++ {
			// 得到当前这个值对应位的值,让其在bucket数组对应位置++
			buckets[nums[j]/int(math.Pow10(i))%10]++
		}
		// 将bucket数组的值累加，可以得到每组数据应该在数组中的区域
		// 累加器
		summ := 0
		for i, v := range buckets {
			if v != 0 {
				buckets[i] += summ
				summ = buckets[i]
			}
		}
		// 将原数组倒序遍历，根据bucket数组的值来填入辅助数组
		for j := len(nums) - 1; j >= 0; j-- {
			// 将对应值填充到辅助数组对应的位置
			help[buckets[nums[j]/int(math.Pow10(i))%10]-1] = nums[j]
			// 位置-1
			buckets[nums[j]/int(math.Pow10(i))%10]--
		}
		// 将辅助数组填充到原数组
		for i := 0; i < len(nums); i++ {
			nums[i] = help[i]
		}
	}

}
func maxBit(nums []int) int {
	max := nums[0]
	for _, val := range nums {
		if val > max {
			max = val
		}
	}
	// 记录位数
	var flag = 0
	for max > 0 {
		max = max / 10
		flag++
	}
	return flag
}
