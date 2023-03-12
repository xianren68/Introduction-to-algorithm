package middle

import "fmt"

// 定义点结构
type drop struct {
	X int
	Y int
}

// 这里y是横轴，x是纵轴

func Zigzag(nums [][]int) {
	// 定义两个数的坐标
	m := &drop{0, 0}
	n := &drop{0, 0}
	for i := 0; m.X >= n.X && m.Y <= n.Y; i++ {
		// 奇数次
		if i%2 == 1 {
			x := n.X
			y := n.Y
			for x <= m.X && y >= m.Y {
				fmt.Println(nums[x][y])
				x++
				y--
			}
			// 偶数次
		} else {
			x := m.X
			y := m.Y
			for x >= n.X && y <= n.Y {
				fmt.Println(nums[x][y])
				x--
				y++
			}
		}
		// 让两个点移动
		if n.Y == len(nums[0])-1 {
			n.X++
		} else {
			n.Y++
		}
		if m.X == len(nums)-1 {
			m.Y++
		} else {
			m.X++
		}

	}
}
