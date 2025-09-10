package main

import (
	"fmt"
	"mylearning/myutil"
)

func main() {
	fmt.Println("Learn go lannguage")
	main1()
	fmt.Println("================================")
	// calling packages
	myutil.PrintMessage("Hello World")
}

func main1() {
	fmt.Println("Carefully learn go lang")
}
