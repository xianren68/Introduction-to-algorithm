package main

import (
	"algorithm/map_test"
	"algorithm/utils"
)

func main() {
	x := [][]int{{1, 3, 7}, {1, 2, 5}, {2, 3, 1}, {2, 6, 3}, {3, 6, 5}, {2, 5, 10}, {1, 4, 5}, {4, 2, 3}}
	f := map_test.TransFromAny(x)
	f.OutPutTel()
	res := map_test.Prime(f)
	utils.PrintEdge(res)

}
