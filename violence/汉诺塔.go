package violence

import "fmt"

func Hanoi(n int, from string, to string, other string) {
	// 只有一个盘子，直接挪
	if n == 1 {
		fmt.Printf("把第个%d盘子从%s挪到%s\n", n, from, to)
		return
	}
	Hanoi(n-1, from, other, to)
	fmt.Printf("把第个%d盘子从挪%s到%s\n", n, from, to)
	Hanoi(n-1, other, to, from)
}
