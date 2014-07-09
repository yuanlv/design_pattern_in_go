/*
	策略模式

	商场打折，根据不同的销售策略方式进行收费
*/

package main

import (
	"fmt"
)

type cashSuper interface {
	AcceptCash(float64) float64
}

type cashNormal struct {
}

func (normal *cashNormal) AcceptCash(money float64) float64 {
	return money
}

type cashRebate struct {
	moneyRebate float64 //打折
}

func (rebate *cashRebate) AcceptCash(money float64) float64 {
	return money *rebate.moneyRebate
}

type CashContext struct {
	cashSuper
}

func NewCashContext(str string) *CashContext {
	cash := new(CashContext)
	switch str {
	case "正常收费":
		cash.cashSuper = &cashNormal{}
	case "打折":
		cash.cashSuper = &cashRebate{0.8}
	}

	return cash
}

func main() {
	var total float64
	context := NewCashContext("正常收费")
	total += context.AcceptCash(1000)	
	fmt.Println("当前收费", total)

	context = NewCashContext("打折")
	total += context.AcceptCash(1000)
	fmt.Println("当前收费", total)

}