package main

import "fmt"

// Public := "this is important"

func custom_variables() {
	var name string = "Prarinam"
	fmt.Println("my name is: ", name)

	var version = 1.1
	fmt.Println("my version is: ", version)

	var currency int = 52436
	fmt.Println("my currency is: ", currency)

	var dimention float64 = 52436.5214
	fmt.Println("my dimention is: ", dimention)

	var decided bool = true
	fmt.Println("my decision is: ", decided)

	const pi float32 = 67.12
	fmt.Println("value of pi is: ", pi)

	// new way to declare variables( not exported)
	person := "Prarinam"
	fmt.Println("my name is: ", person)

	// if variable name start with capital letter 
	// then we can use that from another package also (exported)
	Public := "this is important and exported"
	fmt.Println("my name is: ", Public)
}
