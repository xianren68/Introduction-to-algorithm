package utils

import "fmt"

// 单链表节点
type Node struct {
	Value int
	Next  *Node
}

// 输入一组数转化为单链表
func CreatLink(nums []int) *Node {
	if nums == nil {
		return nil
	}
	// 创建头节点
	head := &Node{Value: nums[0], Next: nil}
	// 遍历的节点
	node := head
	for i := 1; i < len(nums); i++ {
		// 将当前值转化为节点,并连接到node后
		node.Next = &Node{Value: nums[i], Next: nil}
		// node后移一位
		node = node.Next

	}
	return head
}

// 输出链表
func OutLink(head *Node) {
	if head == nil {
		return
	}
	for head != nil {
		fmt.Print(head.Value)
		fmt.Printf(" ")
		head = head.Next
	}

}

// 反转链表
func ReverseLink(head *Node) *Node {
	// 判断临界条件
	if head == nil || head.Next == nil {
		return head
	}
	// 定义两个指针
	cur := head
	pre := head.Next
	for {
		if pre == nil {
			return cur
		}
		// 记录下一个pre的位置
		flag := pre.Next
		if cur == head {
			cur.Next = nil
		}
		pre.Next = cur
		cur = pre
		pre = flag

	}
}
