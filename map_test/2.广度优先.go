package map_test

import (
	"algorithm/utils"
	"fmt"
)

func BFS(node *utils.MapNode) {
	if node == nil {
		return
	}
	// 准备队列
	quene := make([]*utils.MapNode, 0)
	// 准备一个hash表，记录节点是否已经入过队
	hash := make(map[*utils.MapNode]int)
	// 将源节点入队
	quene = append(quene, node)
	hash[node] = 1
	for len(quene) != 0 {
		// 出队
		node = quene[0]
		quene = quene[1:]
		fmt.Println(node.Val)
		// 将邻接结点入队
		for _, val := range node.Nexts {
			// 判断邻接节点是否已经入过队
			if _, ok := hash[val]; !ok {
				quene = append(quene, val)
				hash[val] = 1
			}
		}
	}
}
