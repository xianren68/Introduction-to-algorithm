package basic_advance

import "math"

// 定义每个员工的结构
type Employee struct {
	// 快乐值
	Happy int
	// 他的直接下属
	List []*Employee
}

// 每个子树应该给根节点返回的信息
type HappyInfo struct {
	// 来的最大快乐值
	Attend int
	// 不来的最大快乐值
	Absent int
}

func MaxHappy(boss *Employee) int {
	info := pro(boss)
	return int(math.Max(float64(info.Absent), float64(info.Attend)))
}
func pro(boss *Employee) *HappyInfo {
	// 基层员工没有下级,他的最大快乐值就是他参加的快乐值
	if boss.List == nil {
		return &HappyInfo{boss.Happy, 0}
	}
	// 当前值参加
	attend := boss.Happy
	// 不参加
	absent := 0
	// 参加时获取所有直接下属不参加的快乐值
	for _, val := range boss.List {
		// 员工返回的信息
		resInfo := pro(val)
		// 加入员工不参加的快乐值
		attend += resInfo.Absent
		// 加入最大的快乐值
		absent += int(math.Max(float64(resInfo.Absent), float64(resInfo.Attend)))

	}
	// 返回当前员工的两种可能
	return &HappyInfo{Attend: attend, Absent: absent}

}
