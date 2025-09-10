package main

import "fmt"

func main(){
	age := 25
	height:=5.8234567
	name := "Alice"
	fmt.Println("age: ", age, "height: ", height, "name: ", name)

	fmt.Printf("age is: %d\n", age)
	fmt.Printf("height is: %.3f\n", height)

	fmt.Printf("Type of name is: %T\n", name)
	fmt.Printf("Type of age is: %T\n", age)
	fmt.Printf("Type of hright is: %T\n", height)

	fmt.Printf("Name: %s, Age: %d, Height: %.2f\n", name, age, height)

}