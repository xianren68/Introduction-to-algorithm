package middle

import (
	"errors"
	"time"
)

// 定义动物类
type Animal struct {
	Name string
	Time int64
}

// 猫
type Cat struct {
	Animal
}

// 狗
type Dog struct {
	Animal
}

// 定义猫狗队列
type CandDQueue struct {
	// dog队列
	DogQueue []Dog
	// cat队列
	CatQueue []Cat
}

// 初始化
func InitCandD() *CandDQueue {
	return &CandDQueue{
		DogQueue: make([]Dog, 0),
		CatQueue: make([]Cat, 0),
	}
}

// 判断是狗队列否为空
func (this *CandDQueue) IsDogEmpty() bool {
	if len(this.DogQueue) == 0 {
		return true
	}
	return false
}

// 判断猫队列是否为空
func (this *CandDQueue) IsCatEmpty() bool {
	if len(this.CatQueue) == 0 {
		return true
	}
	return false
}

// 判断整个队列是否为空
func (this *CandDQueue) IsEmpty() bool {
	if this.IsCatEmpty() && this.IsDogEmpty() {
		return true
	}
	return false
}

// 入队
func (this *CandDQueue) Add(a interface{}) {
	// 是狗类型,加入时间戳入队
	if val, ok := interface{}(a).(Dog); ok {
		val.Time = time.Now().Unix()
		this.DogQueue = append(this.DogQueue, val)
	} else {
		val, _ := interface{}(a).(Cat)
		val.Time = time.Now().Unix()
		this.CatQueue = append(this.CatQueue, val)
	}
}

// dog出队
func (this *CandDQueue) PollDog() (Dog, error) {
	if this.IsDogEmpty() {
		return Dog{}, errors.New("队列为空")
	}
	res := this.DogQueue[0]
	this.DogQueue = this.DogQueue[1:]
	return res, nil
}

// cat出队
func (this *CandDQueue) PollCat() (Cat, error) {
	if this.IsCatEmpty() {
		return Cat{}, errors.New("队列为空")
	}
	res := this.CatQueue[0]
	this.CatQueue = this.CatQueue[1:]
	return res, nil
}

// 全部出队
func (this *CandDQueue) PollAll() (interface{}, error) {
	if this.IsEmpty() {
		return Animal{}, errors.New("队列为空")
	}
	if this.IsDogEmpty() {
		res, _ := this.PollCat()
		return res, nil
	}
	if this.IsCatEmpty() {
		res, _ := this.PollDog()
		return res, nil
	}
	if this.CatQueue[0].Time < this.DogQueue[0].Time {
		res := this.CatQueue[0]
		this.CatQueue = this.CatQueue[1:]
		return res, nil
	}
	res := this.DogQueue[0]
	this.DogQueue = this.DogQueue[1:]
	return res, nil

}
