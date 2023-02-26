package utils

import "fmt"

// 图的模板
// 定义图
type MyMap struct {
	// 一个hash表，用来存储每个节点
	NodeHash map[int]*MapNode
	// 用来存储每个边
	EdgeHash map[*Edge]int
}

// 节点
type MapNode struct {
	// 值
	Val int
	// 入度
	In int
	// 出度
	Out int
	// 邻接节点
	Nexts []*MapNode
	// 边
	Edges []*Edge
}

// 边
type Edge struct {
	//权重
	Weight int
	// 来的节点
	Form *MapNode
	// 去的节点
	To *MapNode
}

// 初始化图方法
func (m *MyMap) InitMap() {
	m.NodeHash = make(map[int]*MapNode)
	m.EdgeHash = make(map[*Edge]int)
}

// 初始化节点方法
func (n *MapNode) InitNode(value int) {
	n.Val = value
	n.In = 0
	n.Out = 0
	n.Nexts = make([]*MapNode, 0)
	n.Edges = make([]*Edge, 0)
}

// 初始化边的方法
func (e *Edge) InitEdge(weight int, from, to *MapNode) {
	e.Weight = weight
	e.Form = from
	e.To = to
}

// 模板输出方法
func (t *MyMap) OutPutTel() {
	for _, val := range t.NodeHash {
		fmt.Printf("{值：%#v,入度：%#v,出度:%#v,邻接节点:", val.Val, val.In, val.Out)
		for i, v := range val.Nexts {
			fmt.Printf("（节点值：%#v,", v.Val)
			fmt.Printf("权重：%#v)}", val.Edges[i].Weight)
		}
		fmt.Println()

	}
}

// 输出边数组的方法
func PrintEdge(edges []*Edge) {
	for _, val := range edges {
		fmt.Println("{ from:", val.Form.Val, " to:", val.To.Val, "weight:", val.Weight, "}")
	}
}
