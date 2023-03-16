package middle

func Cur(pre, mid, cu []int, pleft, pright, mleft, mright, cl, cr int) {
	if pleft > pright {
		return
	}
	if pleft == pright {
		cu[cr] = pre[pleft]
		return
	}
	// 先序数组头一定是树的根
	// 找到其在中序数组的位置
	var rIndex int
	for i := mleft; i <= mright; i++ {
		if mid[i] == pre[pleft] {
			rIndex = i
			break
		}
	}
	l := rIndex - mleft
	// 头节点左边全为左子树，右边全为右子树
	// 头节点在末位
	cu[cr] = pre[pleft]
	// 左子树调整
	Cur(pre, mid, cu, pleft+1, pleft+l, mleft, rIndex-1, cl, cl+l-1)
	// 右子树调整
	Cur(pre, mid, cu, pleft+l+1, pright, rIndex+1, mright, cl+l, cr-1)
}
