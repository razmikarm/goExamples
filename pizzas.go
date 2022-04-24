package main

import (
	"fmt"
	"strconv"
	"time"
)

var pizzaNum = 0

func makeDough(stringChan chan string) {

	pizzaNum++
	pizzaName := "Pizza #" + strconv.Itoa(pizzaNum)
	fmt.Println("make Dough for", pizzaName, "and Send for Sauce")
	stringChan <- pizzaName
	time.Sleep(time.Millisecond * 10)
}

func addSauce(stringChan chan string) {

	pizza := <-stringChan
	fmt.Println("Add Sauce and Send", pizza, "for toppings")
	stringChan <- pizza
	time.Sleep(time.Millisecond * 10)
}

func addToppings(stringChan chan string) {

	pizza := <-stringChan
	fmt.Println("Add toppings to", pizza, "and ship")
	time.Sleep(time.Millisecond * 10)
}
