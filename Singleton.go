package main

import "fmt"

//上帝是一个类
type God struct {
	name string
}

var onlyOneGod *God

//每个人都可以得到我的上帝，但上帝只有一个
func GetMyGod() *God {
	if onlyOneGod == nil {
		onlyOneGod = &God{"I am your God"}
	}
	return onlyOneGod
}

func (god *God) Say() {
	fmt.Println(god.name)
}

func main() {
	fmt.Println("design pattern 1 singleton example code")
	GetMyGod().Say()
}
