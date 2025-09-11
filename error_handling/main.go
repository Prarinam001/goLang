package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator must not be zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("Started error handling in functions")

	// ans, err := divide(8, 0)
	// if err != nil{
	// 	fmt.Println(err)
	// }
	// fmt.Println("Division of two number is: ", ans)

	ans, _ := divide(8, 2)
	fmt.Println("Division of two number is: ", ans)

}
