//建造者模式
package main

import "fmt"

//抽象建造者， 生产各种食物的抽象接口
type FoodBuilder interface {
	makeMilk()   //生产牛奶
	makeNoodle() // 生产面条
}

//具体的建造者
type ZhenGongfuFoodBuilder struct {
}

func (this *ZhenGongfuFoodBuilder) makeMilk() {
	fmt.Println("真功夫生产的牛奶")
}

func (this *ZhenGongfuFoodBuilder) makeNoodle() {
	fmt.Println("真功夫生产的面条")
}

type HongZhuangYuanFoodBuilder struct {
}

func (this *HongZhuangYuanFoodBuilder) makeMilk() {
	fmt.Println("宏状元生产的牛奶")
}

func (this *HongZhuangYuanFoodBuilder) makeNoodle() {
	fmt.Println("宏状元生产的面条")
}

//指挥者， 构建使用Builder接口的对象
type FoodDirector struct {
	FoodBuilder
}

func (director *FoodDirector) makeFood() {
	director.makeMilk()
	director.makeNoodle()
}

func main() {
	zhengongfu := &ZhenGongfuFoodBuilder{}
	hongzhuangyuan := &HongZhuangYuanFoodBuilder{}

	// foodDirectors := [2]FoodDirector{
	// 	FoodDirector{zhengongfu},
	// 	FoodDirector{hongzhuangyuan},
	// }

	foodDirectors := []FoodDirector{FoodDirector{zhengongfu}, FoodDirector{hongzhuangyuan}}

	for _, foodDirector := range foodDirectors {
		foodDirector.makeMilk()
		foodDirector.makeNoodle()
	}

}
