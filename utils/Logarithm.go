package utils

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 对数器
func Logar(fn1 func(nums []int, left, right int)) {
	testTimes := 50000 // 测试次数
	maxSize := 1000    //最大测试容量
	maxNum := 1000     //最大测试数据
	equal := true      // 判断每一次比较是否都相等
	// 记录出错的数组
	var input []int
	var target []int
	var my []int
	for i := 0; i < testTimes; i++ {
		// 接收用于测试的两个数组
		arr1, arr2 := randomArr(maxSize, maxNum)
		input = arr1
		// 执行自己的方法
		fn1(arr1, 0, len(arr1)-1)
		// 官方方法
		sort.Ints(arr2)
		// 判断排序结果是否一致
		if !equals(arr1, arr2) {
			target = arr2
			my = arr1
			equal = false
			break
		}
	}
	if equal {
		fmt.Println("完成！")
	} else {
		fmt.Println("失败")
		fmt.Println("输入", input)
		fmt.Println("预期结果")
		fmt.Println(target)
		fmt.Println("我的")
		fmt.Println(my)
	}

}

// 比较方法,判断数组每一位是否相等
func equals(n []int, m []int) bool {
	for i := 0; i < len(n); i++ {
		if n[i] != m[i] {
			return false
		}
	}
	return true
}

// 生成两个相同的随机数组
func randomArr(maxSize int, maxNum int) ([]int, []int) {
	// 获得随机大小
	max := RandomInt(maxSize)
	// 创建一个数组
	arr := make([]int, max)
	// 填充数组
	for k, _ := range arr {
		arr[k] = RandomInt(maxNum)
	}
	// 复制当前切片
	var arr2 []int
	arr2 = append(arr2, arr...)
	return arr, arr2
}

// 随机数函数
func RandomInt(maxSize int) int {
	// 初始化随机数种子
	rand.Seed(time.Now().Unix())
	// 生成一个随机数
	return rand.Intn(maxSize)
}
