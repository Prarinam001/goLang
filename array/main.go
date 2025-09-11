package main

import "fmt"

func main() {
	fmt.Println("Learning array in Golang")

	var name[5] string
	name[0] = "prince"
	name[1] = "akash"

	fmt.Println("Name of persons are: ", name)

	// direct initialization
	var numbers = [8]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers are: ", numbers)

	// check length of arr
	fmt.Println("Length of Numbers array is: ", len(numbers))

	//access particular index
	fmt.Println("value of name at 1 index is: ", name[1])

	var price[5]string
	fmt.Println("Price is: ", price)
	fmt.Printf("Price array is: %q\n", price)
}