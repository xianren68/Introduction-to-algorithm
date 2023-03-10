package middle

func Dedup(nums []int, k int) int {
	// 定义一个hash表，将所有的值加入
	hash := make(map[int]int, len(nums))
	for _, val := range nums {
		hash[val] = 1
	}
	// 数字对
	sum := 0
	for _, val := range nums {
		// 防止重复统计
		if i := hash[val]; i == 2 {
			continue
		}
		// 存在
		if _, ok := hash[val+k]; ok {
			sum++

		}
		hash[val]++
	}
	return sum
}
