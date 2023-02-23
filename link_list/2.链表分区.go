package link_list

import "algorithm/utils"

// 借助辅助空间，
// 1.将链表数据放到数组中partion
// 2.将划分好的数据装回链表
func PartionLinkOne(head *utils.Node, num int) *utils.Node {
	if head == nil || head.Next == nil {
		return head
	}
	node := head
	help := make([]int, 0)
	for node != nil {
		help = append(help, node.Value)
		node = node.Next
	}
	// 划分为三块
	partition(help, num)
	// 遍历填充到链表
	node = head
	for _, val := range help {
		node.Value = val
		node = node.Next
	}
	return head
}

// 荷兰国旗问题，用于将数组分为三部分
func partition(nums []int, num int) {
	// 双指针
	// l代表<区，r代表>区，中间为等于区
	l := 0
	r := len(nums) - 1
	i := 0
	for i <= r {
		// 小于给定值,交换到<区的后一个位置，扩展<区
		if nums[i] < num {
			nums[i], nums[l] = nums[l], nums[i]
			i++
			l++
			// 大于给定值，不发生变化,将当前值交换到>区的前一个位置，拓展>区
		} else if nums[i] > num {
			nums[r], nums[i] = nums[i], nums[r]
			r--
		} else {
			i++
		}

	}
}

// 2. 不借助辅助空间，
// 1.给定六个指针，分别代表<区的头和尾，=区的头和尾，>区的头和尾
// 2.遍历链表，碰到属于哪个区域，就将其连接到那六个指针上
// 3.将指针拼接起来，形成一个新链表
// 连接判断好麻烦能简化吗？
func PartionLinkTwo(head *utils.Node, num int) *utils.Node {
	if head == nil || head.Next == nil {
		return head
	}
	// 六个指针
	var lessStart, lessEnd, equalStart, equalEnd, greatStart, greatEnd *utils.Node
	// 遍历链表
	for head != nil {
		if head.Value < num {
			// 还不存在头节点，
			if lessStart == nil {
				lessStart = &utils.Node{head.Value, nil}
				// 只有一个值，将当前值与头节点连起来
			} else if lessEnd == nil {
				lessEnd = &utils.Node{head.Value, nil}
				lessStart.Next = lessEnd
				// 将当前值与上一个节点相连，并将末位后移
			} else {
				lessEnd.Next = &utils.Node{head.Value, nil}
				lessEnd = lessEnd.Next
			}
		} else if head.Value == num {
			// 还不存在头节点，
			if equalStart == nil {
				equalStart = &utils.Node{head.Value, nil}
				// 只有一个值，将当前值与头节点连起来
			} else if equalEnd == nil {
				equalEnd = &utils.Node{head.Value, nil}
				equalStart.Next = equalEnd
				// 将当前值与上一个节点相连，并将末位后移
			} else {
				equalEnd.Next = &utils.Node{head.Value, nil}
				equalEnd = equalEnd.Next
			}
		} else {
			// 还不存在头节点，
			if greatStart == nil {
				greatStart = &utils.Node{head.Value, nil}
				// 只有一个值，将当前值与头节点连起来
			} else if greatEnd == nil {
				greatEnd = &utils.Node{head.Value, nil}
				greatStart.Next = greatEnd
				// 将当前值与上一个节点相连，并将末位后移
			} else {
				greatEnd.Next = &utils.Node{head.Value, nil}
				greatEnd = greatEnd.Next
			}
		}
		head = head.Next
	}
	if lessStart != nil {
		if lessEnd != nil {
			if equalStart != nil {
				lessEnd.Next = equalStart
				if equalEnd != nil {
					equalEnd.Next = greatStart
				} else {
					equalStart.Next = greatStart
				}
			} else {
				lessEnd.Next = greatStart
			}
		} else {
			if equalStart != nil {
				lessStart.Next = equalStart
				if equalEnd != nil {
					equalEnd.Next = greatStart
				} else {
					equalStart.Next = greatStart
				}
			} else {
				lessStart.Next = greatStart
			}
		}
		return lessStart
	}
	if lessStart == nil && equalStart != nil {
		if equalEnd != nil {
			equalEnd.Next = greatStart
		} else {
			equalStart.Next = greatStart
		}
		return equalStart
	}
	return greatStart

}
