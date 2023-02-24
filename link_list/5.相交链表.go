package link_list

// 相交链表
// 1. 用hash表记录位置，然后对比得到结果
// 2. 不用hash表
func Intersect(head *ListNode, head2 *ListNode) *ListNode {
	// 临界条件
	if head == nil || head2 == nil {
		return nil
	}
	// 判断链表是否有环，若有，接收其入环节点
	loop1 := DetectCycle(head)
	loop2 := DetectCycle(head2)
	// 如果都无环
	if loop1 == nil && loop2 == nil {
		return firstInt(head, head2, nil, nil)
		// 都有环
	} else if loop1 != nil && loop2 != nil {
		if loop1 == loop2 {
			return firstInt(head, head2, loop1, loop2)
		}
		// 转一圈，看能否碰到哦loop2
		node := loop1.Next
		for node != loop1 {
			if node == loop2 {
				return loop1
			}
		}
		// 各自有环，无相交
		return nil
	}
	return nil

}

// 无环情况下第一个相交的节点
func firstInt(head, head2, loop1, loop2 *ListNode) *ListNode {
	//遍历记录长度
	n := 0
	n1 := head
	n2 := head2
	for n1 != loop1 {
		n1 = n1.Next
		n++
	}
	for n2 != loop2 {
		n2 = n2.Next
		n--
	}

	// 根据n的值判断其大小,并将长的记为head,短的为head1,方便遍历
	if n < 0 {
		head, head2 = head2, head
		// 转化为正数
		n = n - 2*n
	}
	// 遍历,先走
	n1 = head
	n2 = head2
	for n != 0 {
		n1 = n1.Next
		n--
	}
	// 同时遍历
	for n1 != n2 {
		n1 = n1.Next
		n2 = n2.Next
	}
	return n1
}
