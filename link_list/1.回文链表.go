package link_list

// 回文链表
import (
	"algorithm/utils"
)

// 01.额外空间，将所有的数据入栈，然后遍历链表和出栈，每个值相比较，不相等则错
func PalOne(head *utils.Node) bool {
	// 判断临界条件
	if head == nil || head.Next == nil {
		return true
	}
	// 只有两个值
	if head.Next.Next == nil {
		return head.Value == head.Next.Value

	}
	// 创建一个数组做为栈
	stack := make([]int, 0)
	// 遍历链表，将每个值入栈
	node := head
	for node != nil {
		stack = append(stack, node.Value)
		node = node.Next
	}
	// 遍历链表和栈对比其值
	// 逆序遍历，模拟出栈的过程
	node = head
	for i := len(stack) - 1; i >= 0; i-- {
		if node.Value != stack[i] {
			return false
		}
		node = node.Next
	}
	return true
}

// 02 额外空间+快慢指针，只用入栈中间值右边的部分
func PalTwo(head *utils.Node) bool {
	// 判断临界条件
	if head == nil || head.Next == nil {
		return true
	}
	// 只有两个值
	if head.Next.Next == nil {
		return head.Value == head.Next.Value
	}
	// 创建一个数组做为栈
	stack := make([]int, 0)
	// 定义两个指针,一个每次走一步一个每次走两步，一个到末位/末位前一位时另一个恰好到中间
	f := head
	l := head
	for f.Next != nil && f.Next.Next != nil {
		f = f.Next.Next
		l = l.Next
	}

	// 将l 到 末位的值入栈，保留l的位置
	for l != nil {
		stack = append(stack, l.Value)
		l = l.Next
	}
	// 出栈与前面部分的值比较
	for i := len(stack) - 1; i >= 0; i-- {
		if head.Value != stack[i] {
			return false
		}
		head = head.Next
	}
	return true

}

// 03 不借助辅助空间，
// 1.先通过双指针找到中间位置和末位值，然后将中间到末位的链表反转
// 2.遍历比较两段链表值，不相等则为false
// 3.将链表复原
func PalThree(head *utils.Node) bool {
	// 判断临界条件
	if head == nil || head.Next == nil {
		return true
	}
	// 只有两个值
	if head.Next.Next == nil {
		return head.Value == head.Next.Value

	}
	// 双指针
	// 定义两个指针,一个每次走一步一个每次走两步，一个到末位/末位前一位时另一个恰好到中间
	f := head
	l := head
	for f.Next != nil && f.Next.Next != nil {
		f = f.Next.Next
		l = l.Next
	}
	// 保存中间位置
	mid := l
	// 从l+1开始反转链表
	l = l.Next
	// 反转链表
	r := utils.ReverseLink(l)
	// 遍历两个链表，比较值
	// 保留两个头节点的位置
	nodeL := head
	nodeR := r
	for nodeR != nil {
		if nodeR.Value != nodeL.Value {
			return false
		}
		nodeR = nodeR.Next
		nodeL = nodeL.Next

	}
	// 还原链表
	//再次反转
	l = utils.ReverseLink(r)
	mid.Next = l
	return true

}
