package basic_advance

// 创建bitArr类型
type bitArr struct {
	// int数组，存储数据
	arr []int
	// 位数组长度
	Length int
}

// 创建位数组
func Create(long int) *bitArr {
	bitarr := new(bitArr)
	// 设置数组长度
	// int类型是32位的
	length := long / 32
	bitarr.arr = make([]int, length)
	bitarr.Length = long
	return bitarr
}

// 获取某一位的值
func (this *bitArr) Get(index int) int {
	// 找到在int数组中的位置
	i := index / 32
	// 具体是第几位
	bitIndex := index % 32
	// 拿到具体值
	// 右移位数，则整个数只有最后一位有值，正好对应要找的值
	b := (this.arr[i] >> bitIndex) & 1
	return b

}

// 设置某一位的值
// 只能0/1
func (this *bitArr) Set(index int, val int) {
	// 找到在int数组中的位置
	i := index / 32
	// 具体是第几位
	bitIndex := index % 32
	if val == 0 {
		// 只有对应位置为0
		this.arr[i] = this.arr[i] & (^(1 << bitIndex))
	} else if val == 1 {
		// 改成1，令一个1左移，得到一个只有对应位为1的数
		// 与数组值或，其对应位置必然为1
		this.arr[i] = this.arr[i] | (1 << bitIndex)

	}

}
