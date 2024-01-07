package main

import "fmt"

func makeAmericano(size float32) {
	fmt.Println("\nMaking an Americano\n------------")
	// make an americano coffee using the coffee machine API
	americano := &CoffeeMachine{}
	beansAmount := (size / 8.0) * 5.0
	americano.startCoffee(beansAmount, 2)
	americano.grindBeans()
	americano.useHotWater(size)
	americano.endCoffee()
	fmt.Println("Americano is ready!")
}

func makeLatte(size float32, foam bool) {
	fmt.Println("\nMaking a latte\n--------------")

	// make a latte coffee using the coffee machine API
	latte := &CoffeeMachine{}
	beansAmount := (size / 8.0) * 5.0
	latte.startCoffee(beansAmount, 4)
	latte.grindBeans()
	latte.useHotWater(size)
	milk := (size / 8.0) * 2
	latte.useMilk(milk)
	latte.doFoam(true)
	latte.endCoffee()
	fmt.Println("Latte is ready!")
}
