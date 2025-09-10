package main

import (
	"fmt"
	"mylearning/myutil"
)
// import custom_variables

func main() {
	fmt.Println("Learn go lannguage")
	main1()
	fmt.Println("================================")
	// calling packages
	myutil.PrintMessage("Hello World")
	custom_variables()
}

func main1() {
	fmt.Println("Carefully learn go lang")
}
