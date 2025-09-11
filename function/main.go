package main

import "fmt"

func SimpleFunction() {
	fmt.Println("I am simple function")
}

// both(a, b) parameter are in same type so int written only one time
func add(a, b int) int {
	return a + b
}

// way - 2
func add_2(a int, b int) (result int) {
	result = a + b
	return
}

func main() {
	fmt.Println("Learning GoLang")
	SimpleFunction()

	ans := add(4, 5)
	fmt.Println("add of two numbers: ", ans)

	ans1 := add_2(4, 5)
	fmt.Println("add of two numbers: ", ans1)
}
