package middle

import (
	"fmt"
	"math"
)

func TopKFrequent(words []string, k int) []string {
	// 定义一个hash表用来存储每个词的词频
	hash := make(map[string]int, len(words))
	for _, val := range words {
		if _, ok := hash[val]; ok {
			hash[val]++
		} else {
			hash[val] = 1
		}
	}
	// 初始化堆
	myheap := initHeap(k)
	// 将它们加入堆中
	for i, val := range hash {
		// 判断堆是否满了，若还没有，直接加入
		if !myheap.isFull() {
			myheap.push(&frequency{val: i, times: val})
		} else {
			// 与第一个值比较
			first := myheap.firstVal()
			if first.times < val || (first.times == val && i < first.val) {
				myheap.pop()
				myheap.push(&frequency{val: i, times: val})
			}
		}
	}
	res := make([]string, k)

	for i := k - 1; i >= 0; i-- {
		fmt.Println(i)
		res[i] = myheap.pop().val
	}
	return res
}

// 定义词-次数的键值堆
type frequency struct {
	val   string
	times int
}

// 定义堆
type heap struct {
	// 用以存储数据的数组
	nums     []*frequency
	heapSize int
}

// 初始化方法
func initHeap(n int) *heap {
	hea := &heap{nums: make([]*frequency, n), heapSize: 0}
	return hea
}

// 加入方法
func (this *heap) push(val *frequency) {
	fmt.Println('x')
	this.insert(val)
	fmt.Println('y')
}
func (this *heap) insert(val *frequency) {
	this.heapSize++
	this.nums[this.heapSize-1] = val
	// 移动加入的元素
	i := this.heapSize - 1
	for i >= 0 {
		if (i-1)/2 >= 0 && (this.nums[i].times < this.nums[(i-1)/2].times || (this.nums[i].times == this.nums[(i-1)/2].times && this.nums[i].val > this.nums[(i-1)/2].val)) {
			this.nums[i], this.nums[(i-1)>>1] = this.nums[(i-1)/2], this.nums[i]
		}
		i = (i - 1) >> 1
	}
}

// 弹出方法
func (this *heap) pop() *frequency {
	res := this.nums[0]
	this.intify()
	return res
}
func (this *heap) intify() {
	this.nums[0] = this.nums[this.heapSize-1]
	this.heapSize--
	i := 0
	for i <= this.heapSize-1 {
		// 只有当前值
		if i<<1+1 > this.heapSize-1 {
			return
		}
		left := this.nums[i<<1+1].times
		right := math.MaxInt
		// 右节点存在
		if i<<1+2 <= this.heapSize-1 {
			right = this.nums[i<<1+2].times
		}
		// 找到较小的,换位
		var min int
		if left > right || (left == right && this.nums[i<<1+2].val > this.nums[i<<1+1].val) {
			min = i<<1 + 2
		} else {
			min = i<<1 + 1
		}
		if this.nums[min].times < this.nums[i].times || (this.nums[min].times == this.nums[i].times && this.nums[min].val > this.nums[i].val) {
			this.nums[min], this.nums[i] = this.nums[i], this.nums[min]
			i = min
		} else {
			return
		}

	}
}

// 判断堆是否满了
func (this *heap) isFull() bool {
	if this.heapSize == len(this.nums) {
		return true
	}
	return false
}

// 堆是否为空
func (this *heap) isEmpty() bool {
	if this.heapSize == 0 {
		return true
	}
	return false
}

// 展示第一个值
func (this *heap) firstVal() *frequency {
	return this.nums[0]
}
