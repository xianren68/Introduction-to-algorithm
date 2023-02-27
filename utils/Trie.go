package utils

// 定义前缀树
// 定义每个节点
type Trie struct {
	// 经过当前节点的次数
	Pass int
	// 以当前节点为末位的个数
	End int
	// 路径的数组
	Nexts []*Trie
}

// 初始化节点
func Constructor() Trie {
	return Trie{Pass: 0, End: 0, Nexts: make([]*Trie, 26)}
}

// 插入字符串
func (this *Trie) Insert(word string) {
	// 将字符串切割为字符数组
	strArr := []byte(word)
	// 遍历字符数组
	node := this
	node.Pass++
	for _, val := range strArr {
		// 找到对应节点的路，判断是否为空
		if node.Nexts[val-'a'] == nil {
			// 为空，新建节点,并连接
			newNode := Constructor()
			node.Nexts[val-'a'] = &newNode
		}
		// 获取对应节点，pass++
		node = node.Nexts[val-'a']
		node.Pass++
	}
	// 整个字符串循环完，最后一个字符end++
	node.End++
}

// 返回字符串是否存在于前缀树中
func (this *Trie) Search(word string) bool {
	// 将其转化为字符数组
	strArr := []byte(word)
	// 遍历
	node := this
	for _, val := range strArr {
		// 没有对应的路，返回false
		if node.Nexts[val-'a'] == nil {
			return false
		}
		// 继续顺着路往下走
		node = node.Nexts[val-'a']
	}
	// 最后一个节点，判断它的end
	if node.End > 0 {
		return true
	}
	return false
}

// 判断是否已有此前缀
func (this *Trie) StartsWith(prefix string) bool {
	// 将其转化为字符数组
	strArr := []byte(prefix)
	// 遍历
	node := this
	for _, val := range strArr {
		// 没有对应的路，返回false
		if node.Nexts[val-'a'] == nil {
			return false
		}
		// 继续顺着路往下走
		node = node.Nexts[val-'a']
	}
	return true
}
