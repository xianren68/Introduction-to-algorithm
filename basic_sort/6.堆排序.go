package basic_sort

func HeapSort(nums []int) {
	// 判断临界条件
	if nums == nil || len(nums) < 2 {
		return
	}
	// 先将给定的数组转化为大根堆
	for i, _ := range nums {
		heapInsert(nums, i)
	}
	// 对堆进行排序
	heapSize := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		// 取出最大值，放到末位,并将最后的节点提前
		nums[0], nums[heapSize] = nums[heapSize], nums[0]
		// 堆的长度-1
		heapSize--
		heapify(nums, 0, heapSize)

	}
}

// 入堆
// 将一个值添加进大根堆，
// 并通过与父节点比较，保持数组结构依然是大根堆
func heapInsert(nums []int, index int) {
	for nums[index] > nums[(index-1)/2] {
		// 交换
		nums[index], nums[(index-1)/2] = nums[(index-1)/2], nums[index]
		// 当前位置变化
		index = (index - 1) / 2
	}
}

// 出堆
// 将根节点弹出，并将最后一个节点放到首位
// 通过与子节点比较，保证数组结构依然是大根堆
func heapify(nums []int, index int, heapSize int) {
	// 记录当前值的左节点
	left := (index * 2) + 1
	// 直到没有字节点,(完全二叉树，没有左节点，自然没有右节点)
	for left <= heapSize {
		// 比较子节点大小
		var max int
		// 右节点存在并且值大于左节点
		if left+1 <= heapSize && nums[left] < nums[left+1] {
			max = left + 1
		} else {
			max = left
		}
		// 子节点大，交换
		if nums[max] > nums[index] {
			nums[max], nums[index] = nums[index], nums[max]
			// 索引变化
			index = max
			left = (index * 2) + 1
		} else {
			// 比子节点都大，终止循环
			return
		}
	}
}
