package main

import (
	"algorithm/middle"
	"fmt"
)

func main() {
	pre := []int{1, 2, 3, 4, 5, 6, 7}
	mid := []int{3, 2, 4, 1, 6, 5, 7}
	cur := make([]int, 7)
	middle.Cur(pre, mid, cur, 0, 6, 0, 6, 0, 6)
	fmt.Println(cur)

}
