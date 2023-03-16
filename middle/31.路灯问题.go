package middle

import "fmt"

// 处理每个例子
func StreetLight(s string) {
	// 开始路灯数记为0。
	sum := 0
	// 遍历所有点
	for i := 0; i < len(s); {
		// 墙，直接继续向下
		if s[i] == 'x' {
			i += 1
		} else { // 不是墙灯的数量直接++
			sum++
			// 到了末位，灯只能安在这里
			if i == len(s)-1 {
				break
			} else {
				// 前面是x,灯只能放在这
				if s[i+1] == 'x' {
					i += 2
				} else { // 前面还有.，灯就放在前一个
					i += 3
				}
			}
		}

	}
	fmt.Println(sum)
}
