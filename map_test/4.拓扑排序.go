package map_test

import (
	"algorithm/utils"
	"fmt"
)

func Topology(graph *utils.MyMap) {
	if graph == nil {
		return
	}
	// 一个hash表，用来修改每个节点的入度
	// 因为直接修改会改变图的结构，得复制出来改
	// 键为顶点，值为剩余得入度
	hash := make(map[*utils.MapNode]int)
	// 一个队列，用来装入度为0的节点
	quene := make([]*utils.MapNode, 0)
	// 找到入口，即唯一的入度为0的点
	for _, val := range graph.NodeHash {
		if val.In == 0 {
			quene = append(quene, val)
		}
		hash[val] = val.In
	}
	// 遍历，直到队列为空
	for len(quene) != 0 {
		// 弹出一个值
		node := quene[0]
		quene = quene[1:]
		// 输出这个值
		fmt.Println(node.Val)
		// 将其邻接节点的入度全部-1，并判断其是否入度为0了
		for _, val := range node.Nexts {
			// 入度-1、
			hash[val]--
			// 入度为0，加入队列
			if hash[val] == 0 {
				quene = append(quene, val)
			}
		}
	}

}
