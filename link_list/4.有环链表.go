package link_list

//
type ListNode struct {
	Val  int
	Next *ListNode
}

// 判断是否有环且返回入环节点
func DetectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	// 定义两个指针,先走对应步数，否则进不了循环
	cur := head.Next
	pre := head.Next.Next
	// 遍历知道两者相遇或pre走到末位/末位的前一个
	for cur != pre {
		if pre.Next == nil || pre.Next.Next == nil {
			return nil
		}
		cur = cur.Next
		pre = pre.Next.Next
	}
	// 相等，让快指针到头节点
	pre = head
	// 遍历，知道它们相等
	for cur != pre {
		cur = cur.Next
		pre = pre.Next
	}
	return pre
}
