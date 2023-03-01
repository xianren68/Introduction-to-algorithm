package basic_advance

import (
	"algorithm/utils"
)

// 定义randomPool
type RandomPool[T comparable] struct {
	// 容器长度
	Size int
	// hash表，(T->int),保证找值时O(1)
	THash map[T]int
	// hash表，(int->T),保证能够随机返回值
	IntHash map[int]T
}

// 初始化方法
func CreateRandomPool[T comparable]() *RandomPool[T] {
	r := new(RandomPool[T])
	r.Size = 0
	r.THash = make(map[T]int)
	r.IntHash = make(map[int]T)
	return r
}

// 插值方法
func (r *RandomPool[T]) Insert(key T) {
	// 判断值是否已存在
	if _, ok := r.THash[key]; !ok {
		// 新增值
		r.THash[key] = r.Size
		r.IntHash[r.Size] = key
		r.Size++
	}

}

// 获取值方法
func (r *RandomPool[T]) GetRandom() T {
	// 获得0-size位置上的随机数
	rand := utils.RandomInt(r.Size)
	return r.IntHash[rand]
}

// 删值方法
func (r *RandomPool[T]) Delete(key T) {
	// 找到对应值的位置
	index := r.THash[key]
	// 将对应位置的值改成最后一个位置的
	r.THash[r.IntHash[r.Size-1]] = index
	r.IntHash[index] = r.IntHash[r.Size-1]
	// 在hash表中删除对应的值
	delete(r.THash, key)
	delete(r.IntHash, r.Size-1)
	// 容器长度-1
	r.Size--

}
