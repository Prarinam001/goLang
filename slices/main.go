package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("numbers slice is: ", numbers)
	fmt.Printf("Number has data type: %T\n", numbers)
	fmt.Println("length of slice is: ", len(numbers))
	fmt.Println("Capacity is: ", cap(numbers))

	numbers = append(numbers, 6, 7, 8, 9, 10, 11, 12, 13)
	fmt.Println("numbers slice is: ", numbers)
	fmt.Println("length of slice is: ", len(numbers))
	fmt.Println("Capacity is: ", cap(numbers))

	// Empty string slice
	name := []string{}
	fmt.Println("Name is: ", name, len(name))

	numbers1 := make([]int, 3, 5)
	numbers1 = append(numbers1, 4)
	numbers1 = append(numbers1, 98)
	fmt.Println("numbers1 slice is: ", numbers1)
	fmt.Println("numbers1 length of slice is: ", len(numbers1))
	fmt.Println("numbers1 Capacity is: ", cap(numbers1))

	// Empty slice using make function
	empty_int_slice := make([]int, 0)
	fmt.Println("empty_int_slice: ", empty_int_slice)
	fmt.Println("empty_int_slice length: ", len(empty_int_slice))

}
