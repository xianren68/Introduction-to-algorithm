package map_test

import "algorithm/utils"

// 转化有向树
func TransFrom(nums [][]int) *utils.MyMap {
	// 初始化一张图
	myMap := new(utils.MyMap)
	myMap.InitMap()
	// 遍历这个矩阵
	for _, val := range nums {
		from := val[0]
		to := val[1]
		weight := val[2]
		// 判断来的节点是否已经存在于当前图中
		if _, ok := myMap.NodeHash[from]; !ok {
			// 若不存在则新建
			fromNode := new(utils.MapNode)
			// 初始化
			fromNode.InitNode(from)
			// 加入
			myMap.NodeHash[from] = fromNode
		}
		// 判断去的节点是否存在图中
		if _, ok := myMap.NodeHash[to]; !ok {
			// 若不存在则新建
			toNode := new(utils.MapNode)
			// 初始化
			toNode.InitNode(to)
			myMap.NodeHash[to] = toNode
		}
		// 获取这两个节点
		fromNode := myMap.NodeHash[from]
		toNode := myMap.NodeHash[to]
		// 将去的节点加入到来节点的邻接节点中
		fromNode.Nexts = append(fromNode.Nexts, toNode)

		// 来的节点出队++
		fromNode.Out++
		// 去的节点入队++
		toNode.In++
		// 将边加入出来的节点中
		edge := new(utils.Edge)
		edge.InitEdge(weight, fromNode, toNode)
		fromNode.Edges = append(fromNode.Edges, edge)
		myMap.EdgeHash[edge] = 1

	}
	return myMap
}

// 无向树
func TransFromAny(nums [][]int) *utils.MyMap {
	// 初始化一张图
	myMap := new(utils.MyMap)
	myMap.InitMap()
	// 遍历这个矩阵
	for _, val := range nums {
		from := val[0]
		to := val[1]
		weight := val[2]
		// 判断来的节点是否已经存在于当前图中
		if _, ok := myMap.NodeHash[from]; !ok {
			// 若不存在则新建
			fromNode := new(utils.MapNode)
			// 初始化
			fromNode.InitNode(from)
			// 加入
			myMap.NodeHash[from] = fromNode
		}
		// 判断去的节点是否存在图中
		if _, ok := myMap.NodeHash[to]; !ok {
			// 若不存在则新建
			toNode := new(utils.MapNode)
			// 初始化
			toNode.InitNode(to)
			myMap.NodeHash[to] = toNode
		}
		// 获取这两个节点
		fromNode := myMap.NodeHash[from]
		toNode := myMap.NodeHash[to]
		// 将两个节点分别加入到它们的邻接节点
		fromNode.Nexts = append(fromNode.Nexts, toNode)
		toNode.Nexts = append(toNode.Nexts, fromNode)
		// 两个节点出入队都++
		fromNode.Out++
		toNode.Out++
		// 去的节点入队++
		toNode.In++
		fromNode.In++
		// 将边加入出来两个节点中
		edge := new(utils.Edge)
		edge.InitEdge(weight, fromNode, toNode)
		fromNode.Edges = append(fromNode.Edges, edge)
		edge = new(utils.Edge)
		edge.InitEdge(weight, toNode, fromNode)
		toNode.Edges = append(toNode.Edges, edge)
		myMap.EdgeHash[edge] = 1
	}
	return myMap
}
