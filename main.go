package main

import (
	basic_sort "algorithm/basic_sort"
	"algorithm/utils"
	"fmt"
)

func main() {
	a := []int{1, 2, 5, 8, 6, 7, 3}
	fmt.Println(basic_sort.InsertSort(a))
	utils.Logar(basic_sort.Bubble)

}
