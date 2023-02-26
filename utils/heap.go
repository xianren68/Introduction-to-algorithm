package utils

// 将一个数组转化为堆
// 定义一个比较器方法
func HeapSort[T any](nums []T, fn func(a, b T) bool) {
	// 判断临界条件
	if nums == nil || len(nums) < 2 {
		return
	}
	// 先将给定的数组转化为大根堆
	for i, _ := range nums {
		HeapInsert(nums, i, fn)
	}
	// 对堆进行排序
	heapSize := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		// 取出最大/小值，放到末位,并将最后的节点提前
		nums[0], nums[heapSize] = nums[heapSize], nums[0]
		// 堆的长度-1
		heapSize--
		Heapify(nums, 0, heapSize, fn)

	}
}

// 入堆
// 将一个值添加进堆，
// 并通过与父节点比较，保持数组结构依然是堆
func HeapInsert[T any](nums []T, index int, fn func(a, b T) bool) {
	for fn(nums[(index-1)/2], nums[index]) {
		// 交换
		nums[index], nums[(index-1)/2] = nums[(index-1)/2], nums[index]
		// 当前位置变化
		index = (index - 1) / 2
	}
}

// 出堆
// 将根节点弹出，并将最后一个节点放到首位
// 通过与子节点比较，保证数组结构依然是大根堆
func Heapify[T any](nums []T, index int, heapSize int, fn func(a, b T) bool) {
	// 记录当前值的左节点
	left := (index * 2) + 1
	// 直到没有左节点,(完全二叉树，没有左节点，自然没有右节点)
	for left <= heapSize {
		// 比较子节点大小
		var max int
		// 右节点存在并且值大/小于左节点
		if left+1 <= heapSize && fn(nums[left], nums[left+1]) {
			max = left + 1
		} else {
			max = left
		}
		// 子节点大/小，交换
		if fn(nums[index], nums[max]) {
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
