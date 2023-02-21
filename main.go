package main

import (
	"algorithm/basic_sort"
	"fmt"
)

func main() {
	x := []int{1, 4, 3, 2, 5, 6, 7, 3, 6, 24, 55, 12, 77}

	basic_sort.QuickSort(x, 0, len(x)-1)
	fmt.Println(x)
	// utils.Logar(basic_sort.QuickSort)

}
