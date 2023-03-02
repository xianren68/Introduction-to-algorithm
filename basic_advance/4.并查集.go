package basic_advance

// 定义并查集单个节点的结构
type DSUNode[T comparable] struct {
	// 节点里的值
	Val  T
	Next *DSUNode[T]
	// 以当前节点为头的集合的长度
	Size int
}

// 	定义并查集结构
type DSU[T comparable] struct {
	// 存储值和集合的映射
	Hash map[T]*DSUNode[T]
}

// 初始化方法将给定序列的值全部转化为集合
func InitDSU[T comparable](list []T) *DSU[T] {
	// 新建DSU
	dsu := new(DSU[T])
	dsu.Hash = make(map[T]*DSUNode[T])
	for _, val := range list {
		t := &DSUNode[T]{Val: val, Next: nil, Size: 1}
		dsu.Hash[val] = t
	}
	return dsu
}

// 往上找到头的方法
func (this *DSU[T]) findHead(val T) *DSUNode[T] {
	// 根据hash表找出对应的节点
	ele := this.Hash[val]
	// 准备一个数组，装这些沿途经过的节点
	stack := make([]*DSUNode[T], 0)
	// 遍历到头
	for ele.Next != nil {
		stack = append(stack, ele)
		ele = ele.Next
	}
	head := ele
	// 遍历将所有的节点都直接连接到头
	for _, val := range stack {
		val.Next = head
		val.Size = 0
	}
	return head
}

// 判断是否在同一个集合中
func (this *DSU[T]) IsSameSet(v1, v2 T) bool {
	_, ok1 := this.Hash[v1]
	_, ok2 := this.Hash[v2]
	// 如果都存在则返回它们的头节点是否相等
	if ok1 && ok2 {
		return this.findHead(v1) == this.findHead(v2)
	}
	return false
}

// 合并两个集合
func (this *DSU[T]) UnionSet(v1, v2 T) {
	_, ok1 := this.Hash[v1]
	_, ok2 := this.Hash[v2]
	// 如果都存在则找到它们的头节点，判断长度
	if ok1 && ok2 {
		head1 := this.findHead(v1)
		head2 := this.findHead(v2)
		// 是同一条链
		if head1 == head2 {
			return
		}
		// 合并
		if head1.Size >= head2.Size {
			head2.Next = head1
			head1.Size += head2.Size
			head2.Size = 0
		}
	}
}
