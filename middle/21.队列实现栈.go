package middle

type MyStack struct {
	Queue []int
}

func Constructor() MyStack {
	return MyStack{Queue: make([]int, 0)}

}

func (this *MyStack) Push(x int) {
	// 直接加入队列
	this.Queue = append(this.Queue, x)
}

func (this *MyStack) Pop() int {
	// 将队列中除最后一个元素外的所有元素取出重新加入
	// 队列的长度
	size := len(this.Queue)
	for i := 0; i < size-1; i++ {
		res := this.Queue[0]
		this.Queue = this.Queue[1:]
		this.Queue = append(this.Queue, res)
	}
	// 返回队列最后一个元素
	res := this.Queue[0]
	this.Queue = this.Queue[1:]
	return res
}

func (this *MyStack) Top() int {
	// 将队列中除最后一个元素外的所有元素取出重新加入
	// 队列的长度
	size := len(this.Queue)
	var res int
	for i := 0; i < size; i++ {
		res = this.Queue[0]
		this.Queue = this.Queue[1:]
		this.Queue = append(this.Queue, res)
	}
	return res
}

func (this *MyStack) Empty() bool {
	if len(this.Queue) == 0 {
		return true
	}
	return false
}
