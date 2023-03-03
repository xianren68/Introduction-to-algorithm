package basic_advance

func MaxSlidingWindow(nums []int, k int) []int {
	// 定义一个数组用作返回值
	res := make([]int, 0)
	// 定义一个数组用作双向队列
	queue := make([]int, 0)
	// 先将前几个值加入
	for i := 0; i < k; i++ {
		// 右边界移动,加入值
		queue = push(queue, nums[i])
		if i == len(nums)-1 {
			break
		}
	}
	// 窗口左边界
	L := 0
	// 取得第一个值
	res = append(res, queue[0])
	// 遍历移动窗口
	for i := k; i < len(nums); i++ {
		//先移动右边
		queue = push(queue, nums[i])
		// 再移动左边
		queue = pop(queue, nums[L])
		// 窗口移动
		L++
		// 拿到最大值，并加入数组
		res = append(res, queue[0])
	}
	return res

}
func push(queue []int, val int) []int {
	if len(queue) == 0 {
		queue = append(queue, val)
		return queue
	}
	for len(queue) > 0 && queue[len(queue)-1] < val {
		// 队列长度减小
		queue = queue[:len(queue)-1]
	}
	queue = append(queue, val)
	return queue
}
func pop(queue []int, val int) []int {
	if queue[0] == val {
		// 队列弹出一个值
		queue = queue[1:]
	}
	return queue
}
