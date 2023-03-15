package middle

import (
	"fmt"
	"sort"
	"strings"
)

// 定义前缀树每个节点的类型
type trieNode struct {
	Val  string
	Hash map[string]*trieNode
}

// 定义一个结构体记录层数
type stor struct {
	Val    *trieNode
	Storey int
}

// 初始化节点
func InitTrie(val string) *trieNode {
	return &trieNode{Val: val, Hash: make(map[string]*trieNode)}
}

// 将字符串添加到前缀树
func (this *trieNode) addString(arr []string) {
	// 遍历字符串数组
	trie := this
	for _, val := range arr {
		// 不存在，新建节点,并连接到后面
		if _, ok := trie.Hash[val]; !ok {
			newNode := InitTrie(val)
			trie.Hash[val] = newNode
		}
		trie = trie.Hash[val]
	}
}

// 打印目录
func PrintDir(s []string) {
	// 初始化前缀树
	trie := InitTrie("")
	// 遍历每个字符串，将它们以\分割
	for _, val := range s {
		arr := strings.Split(val, "\\")
		if arr[len(arr)-1] == "" {
			arr = arr[:len(arr)-1]
		}
		trie.addString(arr)
	}
	// 深度优先遍历
	// 定义一个数组作为栈
	stack := make([]*stor, 0)
	// 将第一个节点入栈,并将层数记为-1
	stack = append(stack, &stor{Val: trie, Storey: -1})
	// 遍历输出
	for len(stack) != 0 {
		// 节点出栈
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		var list = make([]*trieNode, len(node.Val.Hash))
		// 将后续节点全部取出放到数组中
		i := 0
		for _, val := range node.Val.Hash {
			list[i] = val
			i++
		}
		// 按字典序排序
		sort.Slice(list, func(i, j int) bool {
			// 将字典序小的靠后
			return list[i].Val > list[j].Val
		})

		// 将前缀树后续节点全部入栈
		for _, val := range list {
			// 层数+1
			stack = append(stack, &stor{Val: val, Storey: node.Storey + 1})
		}
		// 打印节点值
		// 前缀树根节点不用打印
		if node.Storey != -1 {
			s := strings.Repeat("  ", node.Storey)
			fmt.Println(s + node.Val.Val)
		}
	}

}
