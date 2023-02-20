package bitwise

import "fmt"

// 根据异或特性，让一个零与这个数组所有值异或，得到的就是所需的

func QuestionOne(nums []int) int {
	// 定义0
	num := 0
	// 遍历异或每个数据
	for _, val := range nums {
		num ^= val
	}
	return num
}

// 与上题一致，最后的到两个数异或的结果xor
// 两个数不相等，则肯定有一位一个为0,一个为1,xor的这一位为1
// 将结果转化，得到xor最右边的1(只有这一位为1，其余全为0),right,将这个数复制为rightOne
// 将rightOne与数组中所有能和right与运算为0的数进行异或（因为两个数在这一位一个为0,一个为1，与运算为0
// 的值中必然有一个值为想要的值）得到一个异或值
// 最后将得到的值与right异或，得到一个值,然后再将值与xor异或，得到另一个值
func QuestionTwo(nums []int) {
	// 定义0
	num := 0
	// 遍历异或每个数据
	for _, val := range nums {
		num ^= val
	}
	// 最右边的1
	right := num & (^num + 1)
	rightOne := right
	for _, val := range nums {
		if right&val == 0 {
			rightOne ^= val
		}
	}
	fmt.Println(right^rightOne, right^rightOne^num)

}
