package link_list

// 定义这个特殊的节点
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 1.利用hash表，将节点复制
// 1. 将链表源节点作为键，将复制节点作为值
// 2. 遍历链表，每个节点对应由一个复制节点，通过hash表可以得到
// 3. 原节点指针的指向的值，也有对应的映射出的复制节点
// 4. 通过hash表得到复制值，再通过hash表得到其指针指向的复制值，连接起来即可
func copyRandomListOne(head *Node) *Node {
	// 临界条件
	if head == nil {
		return nil
	}
	// 建立hash表
	hash := make(map[*Node]*Node, 0)
	// 遍历链表，将值填入
	node := head
	// 记录新链表的头节点
	var newHead *Node
	for node != nil {
		// 复制节点
		// 建立映射关系
		hash[node] = &Node{Val: node.Val, Next: node.Next}
		node = node.Next
	}
	// 遍历，将随机指针也复制
	node = head
	for node != nil {
		copy := hash[node]
		// 通过映射关系找到对应的值并建立联系
		copy.Next = hash[node.Next]
		copy.Random = hash[node.Random]
		if node == head {
			newHead = copy
		}
		node = node.Next
	}
	return newHead
}

// 2.不使用额外空间
// 1.在每个节点后创建副本与它连接
// 2.遍历整个链表，每次走两格，走过的值均为原链表的值，
// 3.判断值的random指针情况，若为空，则副本也为空，若不为空，它指向某个节点，副本就指向那个节点后它的副本，完成random指针的复制
// 4.通过遍历分离链表
func copyRandomListTwo(head *Node) *Node {
	// 临界条件
	if head == nil {
		return nil
	}
	node := head
	for node != nil {
		// 创建副本，插入链表
		copy := &Node{Val: node.Val}
		copy.Next = node.Next
		node.Next = copy
		// 位置移动
		node = copy.Next
	}
	// 连接随机指针
	node = head
	// 每次走两格，走过的节点都是原链表上的
	for node != nil {
		if node.Random == nil {
			// 随机指针为空，副本也为空
			node.Next.Random = nil
		} else {
			// 随机指针节点
			random := node.Random
			// 它的下一个为它的副本，直接连接两个副本节点
			node.Next.Random = random.Next
		}
		node = node.Next.Next
	}
	// 将副本分离
	// 记录头节点
	newHead := head.Next
	copynode := newHead
	node = head
	for node != nil {
		if node != head {
			copynode.Next = node.Next
			copynode = copynode.Next
		}
		// 跳过副本
		node.Next = node.Next.Next
		node = node.Next

	}
	return newHead
}
