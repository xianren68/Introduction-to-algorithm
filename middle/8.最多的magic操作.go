package middle

import "sort"

func Magic(num1, num2 []int) int {
	if len(num1) == 0 || len(num2) == 0 {
		return 0
	}
	// 求出两个数组的累加和
	suNum1 := 0
	for _, val := range num1 {
		suNum1 += val
	}
	suNum2 := 0
	for _, val := range num2 {
		suNum2 += val
	}
	avgNum1 := avg(suNum1, len(num1))
	avgNum2 := avg(suNum2, len(num2))
	var more, less []int
	var suMo, suLe int
	// 将两个数组重定向
	if avgNum1 == avgNum2 {
		return 0
	} else if avgNum1 > avgNum2 {
		more = num1
		suMo = suNum1
		less = num2
		suLe = suNum2
	} else {
		less = num1
		suLe = suNum1
		more = num2
		suMo = suNum2
	}
	// 为了去重将小数组装入到hash表中
	leHash := make(map[int]int, len(less))
	for _, val := range less {
		leHash[val] = 1
	}
	// 将较大的数组排序
	sort.Ints(more)
	// 记录长度
	le := len(leHash)
	mo := len(more)
	// 操作的次数
	var x = 0
	for _, val := range more {
		if val > int(avg(suLe, le))+1 && val < int(avg(suMo, mo))+1 {
			if _, ok := leHash[val]; !ok {
				suLe += val
				le++
				suMo -= val
				mo--
				x++
			}
		}
	}
	return x

}

// 平均值函数
func avg(sum, n int) float64 {
	return float64(sum) / float64(n)
}
