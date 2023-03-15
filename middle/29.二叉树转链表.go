package middle

import "fmt"

func tranLink(root *TreeNode) *TreeNode {
	head, _ := f(root)
	for head != nil {
		fmt.Println(head.Val)
		head = head.RightNode
	}
	return head
}
func f(root *TreeNode) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}
	// 左节点返回
	lhead, ltail := f(root.LeftNode)
	// 右节点返回
	rhead, rtail := f(root.RightNode)
	head, tail := root, root
	if ltail != nil {
		ltail.RightNode = root
		root.LeftNode = ltail
		head = lhead
	}
	if rhead != nil {
		rhead.LeftNode = root
		root.RightNode = rhead
		tail = rtail
	}
	return head, tail
}
