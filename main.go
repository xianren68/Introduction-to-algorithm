package main

import (
	"algorithm/basic_sort"
	"fmt"
)

func main() {
	x := []int{7, 3, 5, 9, 3, 7, 5, 11, 22, 57, 11, 112, 67}
	basic_sort.BucketSort(x, 10)
	fmt.Println(x)

}
