package nqueen

import (
	"fmt"
)

// 创建一个二维数组表示棋盘
var chessboard = [8][8]int{}

// 每一列，每个对角线，反对角线
var col, dg, udg [20]bool

func dfs(i int) {
	// 已经到了棋盘外了，每一行都走过了，输出这个解法
	if i == 8 {
		for _, val := range chessboard {
			fmt.Println(val)
		}
		fmt.Println("********")
	}
	// 找出当前行符合规范的位置
	for j := 0; j < 8; j++ {
		// 判断是否符合规范
		if !col[j] && !dg[j-i+8] && !udg[j+i] {
			// 将对应列，对角线设置为true
			col[j], dg[j-i+8], udg[j+i] = true, true, true
			// 填入当前值
			chessboard[i][j] = 1
			// 进入下一行
			dfs(i + 1)
			// 回溯
			// 将值返回，判断本行的其他可能节点
			col[j], dg[j-i+8], udg[j+i] = false, false, false
			chessboard[i][j] = 0
		}
	}
}
func One() {
	dfs(0)
}
