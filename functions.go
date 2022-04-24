package main

import "fmt"

func factorial(num int) int {
	if num < 2 {
		return 1
	}
	return num * factorial(num-1)
}

func printOne() { fmt.Println(1) }
func printTwo() { fmt.Println(2) }

func safeDiv(num1, num2 int) int {

	defer func() {
		errorMsg := recover()
		fmt.Println(errorMsg)
	}()

	// panic("PANIC")

	return num1 / num2

}
