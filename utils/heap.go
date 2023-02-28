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

// 并通过与父节点比较，保持数组结构依然是堆
func HeapInsert[T any](nums []T, index int, fn func(a, b T) bool) {
	for fn(nums[(index-1)/2], nums[index]) {
		// 交换
		nums[index], nums[(index-1)/2] = nums[(index-1)/2], nums[index]
		// 当前位置变化
		index = (index - 1) / 2
	}
}

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

// 改变堆中某个位置的值，让其依然保持结构
func InsertHeapify[T any](nums []T, t T, trans T, fn func(a, b T) bool) {
	var index int
	for i, val := range nums {
		if &val == &t {
			index = i
		}
	}
	// 修改对应位置的值
	nums[index] = trans
	// 判断变化后的大小
	if fn(nums[(index/2)-1], nums[index]) {
		HeapInsert(nums, index, fn)
	} else {
		Heapify(nums, index, len(nums), fn)
	}
}

// 弹出第一个值
func HeapPop[T any](nums []T, fn func(a, b T) bool) T {
	res := nums[0]
	nums[0] = nums[len(nums)-1]
	nums = nums[:len(nums)-1]
	Heapify(nums, 0, len(nums)-1, fn)
	return res
}

// 入堆
func HeapPush[T any](nums []T, n T, fn func(a, b T) bool) {
	nums = append(nums, n)
	HeapInsert(nums, len(nums)-1, fn)
}
