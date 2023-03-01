package main

import (
	"algorithm/basic_advance"
	"fmt"
)

func main() {
	bitarr := basic_advance.Create(128)
	bitarr.Set(2, 1)
	fmt.Println(bitarr.Get(2))
	bitarr.Set(1, 1)
	fmt.Println(bitarr.Get(10))
	fmt.Println(bitarr.Length)

}
