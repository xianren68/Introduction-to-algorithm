package basic_advance

func maxAreaOfIsland(grid [][]int) int {
	if grid == nil {
		return 0
	}
	max := 0
	// 遍历整个二维数组
	for i, _ := range grid {
		for j, _ := range grid[i] {
			if grid[i][j] == 1 {
				res := render(grid, i, j)
				if res > max {
					max = res
				}
			}
		}
	}
	return max

}
func render(grid [][]int, i, j int) int {
	if (i < 0 || i >= len(grid)) || (j < 0 || j >= len(grid[0]) || grid[i][j] != 1) {
		return 0
	}
	res := 1
	// 将已经到过的1改成2
	grid[i][j] = 2
	// 继续向四个方向渲染
	res += render(grid, i-1, j)
	res += render(grid, i+1, j)
	res += render(grid, i, j+1)
	res += render(grid, i, j-1)
	return res
}
