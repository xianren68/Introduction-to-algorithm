package main

import (
	"algorithm/link_list"
	"algorithm/utils"
)

func main() {
	x := []int{1, 4, 5, 2, 3, 6, 7}
	head := utils.CreatLink(x)
	head = link_list.PartionLinkTwo(head, 2)
	utils.OutLink(head)

}
