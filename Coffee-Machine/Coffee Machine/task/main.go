package main

import (
	"fmt"
)

func min(values ...int) int {
	now := int(1e9)
	for _, x := range values {
		if x < now {
			now = x
		}
	}
	return now
}

func max(values ...int) int {
	now := int(-1e9)
	for _, x := range values {
		if x > now {
			now = x
		}
	}
	return now
}

type Coffee struct {
	needWater  int
	needMilk   int
	needCoffee int
	cost       int
}

func (c Coffee) NeedWater() int {
	return c.needWater
}

func (c Coffee) NeedMilk() int {
	return c.needMilk
}

func (c Coffee) NeedCoffee() int {
	return c.needCoffee
}

func (c Coffee) Cost() int {
	return c.cost
}

func NewCoffee(needWater int, needMilk int, needCoffee int, cost int) *Coffee {
	return &Coffee{needWater: needWater, needMilk: needMilk, needCoffee: needCoffee, cost: cost}
}

func main() {
	nowMoney, nowWater, nowMilk, nowCoffee, nowCups := 550, 400, 540, 120, 9
	espresso := NewCoffee(250, 0, 16, 4)
	latte := NewCoffee(350, 75, 20, 7)
	cappuccino := NewCoffee(200, 100, 12, 6)

	//getMaxCanMake := func(c Coffee) int {
	//	var qWater, qMilk, qCoffee int
	//	qWater = nowWater / c.NeedWater()
	//	if c.NeedMilk() > 0 {
	//		qMilk = nowMilk / c.NeedMilk()
	//	} else {
	//		qMilk = int(1e9)
	//	}
	//	qCoffee = nowCoffee / c.NeedCoffee()
	//	return min(qWater, qMilk, qCoffee, nowCups)
	//}

	printState := func() {
		fmt.Println("The coffee machine has:")
		fmt.Println(nowWater, "ml of water")
		fmt.Println(nowMilk, "ml of milk")
		fmt.Println(nowCoffee, "g of coffee beans")
		fmt.Println(nowCups, "disposable cups")
		fmt.Printf("$%d of money\n", nowMoney)
	}

	makeCoffee := func(c Coffee, args ...int) {
		nCups := 1
		if len(args) > 0 {
			nCups = args[0]
		}
		fmt.Println("I have enough resources, making you a coffee!")
		nowWater -= c.NeedWater() * nCups
		nowMilk -= c.NeedMilk() * nCups
		nowCoffee -= c.NeedCoffee() * nCups
		nowMoney += c.Cost() * nCups
		nowCups -= 1
	}

	canMake := func(c Coffee) bool {
		if nowWater < espresso.NeedWater() {
			fmt.Println("Sorry, not enough water!")
		} else if nowMilk < espresso.NeedMilk() {
			fmt.Println("Sorry, not enough milk!")
		} else if nowCoffee < espresso.NeedCoffee() {
			fmt.Println("Sorry, not enough coffee beans!")
		} else if nowCups < 1 {
			fmt.Println("Sorry, not enough disposable cups!")
		} else {
			return true
		}
		return false
	}

	makeCoffeeWrapper := func(c Coffee) {
		if canMake(c) {
			makeCoffee(c)
		}
	}

	buy := func() {
		fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino: ")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			makeCoffeeWrapper(*espresso)
		case 2:
			makeCoffeeWrapper(*latte)
		case 3:
			makeCoffeeWrapper(*cappuccino)
		default:
			fmt.Println("Invalid Option")
		}
	}

	fill := func() {
		var addWater, addMilk, addCoffee, addCups int
		fmt.Println("Write how many ml of water you want to add:")
		fmt.Scan(&addWater)
		fmt.Println("Write how many ml of milk you want to add:")
		fmt.Scan(&addMilk)
		fmt.Println("Write how many grams of coffee beans you want to add:")
		fmt.Scan(&addCoffee)
		fmt.Println("Write how many disposable cups you want to add:")
		fmt.Scan(&addCups)
		nowWater += max(0, addWater)
		nowMilk += max(0, addMilk)
		nowCoffee += max(0, addCoffee)
		nowCups += max(0, addCups)
	}

	take := func() {
		fmt.Printf("I gave you $%v\n", nowMoney)
		nowMoney = 0
	}

	prompt := func() {
		running := true
		for running {
			fmt.Println("Write action (buy, fill, take, remaining, exit):")
			var action string
			fmt.Scan(&action)

			switch action {
			case "buy":
				buy()
			case "fill":
				fill()
			case "take":
				take()
			case "remaining":
				printState()
			case "exit":
				running = false
			default:
				fmt.Println("Invalid Option")
			}
			println()
		}
	}

	prompt()
}
