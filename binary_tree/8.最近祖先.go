package binary_tree

// 前提：p,q都在子树中
// 1.hash表
func lowestCommonAncestorOne(root, p, q *TreeNode) *TreeNode {
	// 建立一个hash表
	hash := make(map[*TreeNode]*TreeNode, 0)
	// 令头节点进入hash表，并让它的父节点为空
	hash[root] = nil
	// 遍历子树，将每个节点与其父节点加入hash表
	pushHash(hash, root)
	// 抽离出一个节点的祖先链
	hash1 := make(map[*TreeNode]*TreeNode, 0)
	for q != nil {
		// go中没有set结构，这里的root只用作占位，表示这个值存在
		hash1[q] = root
		// 顺着祖先链回溯
		q = hash[q]
	}
	// 通过回溯另一个值来对比
	for p != nil {
		// 判断它是不是存在于另一个节点的祖先链上
		if _, ok := hash1[p]; ok {
			// 存在，返回即可
			return p
		}
		// 继续回溯
		p = hash[p]

	}
	return root

}

// 将值填充进hash表
func pushHash(hash map[*TreeNode]*TreeNode, root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		hash[root.Left] = root
		pushHash(hash, root.Left)
	}
	if root.Right != nil {
		hash[root.Right] = root
		pushHash(hash, root.Right)
	}
}

// 2.递归
func lowestCommonAncestorTwo(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	l := lowestCommonAncestorTwo(root.Left, p, q)
	r := lowestCommonAncestorTwo(root.Right, p, q)
	// 二者都不为空，则说明p,q在当前子树的两侧，则它们的公共祖先必然为当前节点
	if r != nil && l != nil {
		return root
	}
	// 都为空，直接返回nil,当前子树没有这两个值
	if r == nil && l == nil {

	}
	// 返回不为空的一个，它就是公共祖先
	if r != nil {
		return r
	}
	return l

}
