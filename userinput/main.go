package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hey, what's your name")
	// var name string

	// Scan() method reads until the first whitespace character
	// fmt.Scan(&name)
	// fmt.Println("Hello Mr. ", name)

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello, Mr: ", name)
}
