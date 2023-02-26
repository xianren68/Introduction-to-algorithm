package map_test

import (
	"algorithm/utils"
	"fmt"
)

func DFS(node *utils.MapNode) {
	if node == nil {
		return
	}
	// 准备一个栈
	stack := make([]*utils.MapNode, 0)
	// 准备一个hash表
	hash := make(map[*utils.MapNode]int)
	// 将源节点入栈
	stack = append(stack, node)
	hash[node] = 1
	for len(stack) != 0 {
		// 弹出节点
		node = stack[len(stack)-1]
		fmt.Println(node.Val)
		stack = stack[:len(stack)-1]
		for _, val := range node.Nexts {
			// 判断邻接节点是否已经入过栈
			if _, ok := hash[val]; !ok {
				stack = append(stack, val)
				hash[val] = 1
			}
		}
	}
}
