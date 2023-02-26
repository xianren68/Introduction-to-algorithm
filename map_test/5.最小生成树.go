package map_test

import (
	"algorithm/utils"
	"fmt"
)

// 最小生成树Kruskal算法
func Kruskal(graph *utils.MyMap) {
	if graph == nil {
		return
	}
	// 建立一个新图，作为最小生成树
	newMap := new(utils.MyMap)
	newMap.InitMap()
	// 将所有边复制到一个数组中，转化为小根堆
	edges := make([]*utils.Edge, 0)
	for key, _ := range graph.EdgeHash {
		edges = append(edges, key)
	}

	// 转化
	utils.HeapSort(edges, coml)
	for _, val := range edges {
		fmt.Println(val.Weight)
	}

}

// 定义比较器方法
func coml(a, b *utils.Edge) bool {
	return b.Weight-a.Weight < 0
}

func Prime(graph *utils.MyMap) []*utils.Edge {
	if graph == nil {
		return nil
	}
	// 一个数组，用来返回最小生成树的所有边
	res := make([]*utils.Edge, 0)

	// 准备一个堆
	heap := make([]*utils.Edge, 0)
	// 准备一个hash表，用来避免节点重复加入
	hash := make(map[*utils.MapNode]int)
	// 加入一个节点作为开始节点
	var node *utils.MapNode
	for _, val := range graph.NodeHash {
		node = val
		break
	}
	// 节点存入hash表
	hash[node] = 1
	for i := 1; i < len(graph.NodeHash); i++ {
		// 将节点的所有边入堆
		for _, val := range node.Edges {
			heap = append(heap, val)
			utils.HeapInsert(heap, len(heap)-1, coml)

		}
		// 出一个最小的
		min := heap[0]
		heap[0] = heap[len(heap)-1]
		heap = heap[:len(heap)-1]
		// 保持小根堆
		utils.Heapify(heap, 0, len(heap)-1, coml)
		node = min.To
		// 若node不存在，则加入
		if _, ok := hash[node]; !ok {
			hash[node] = 1
			// 加入返回的数组中
			res = append(res, min)
			continue
		}
		// 遍历找到一个不存在的点
		for len(heap) > 0 {
			if _, ok := hash[node]; !ok {
				break
			}
			// 出一个最小的
			min = heap[0]
			heap[0] = heap[len(heap)-1]
			heap = heap[:len(heap)-1]
			utils.Heapify(heap, 0, len(heap)-1, coml)
			node = min.To
		}
		hash[node] = 1
		// 加入返回的数组中
		res = append(res, min)
	}
	return res
}
