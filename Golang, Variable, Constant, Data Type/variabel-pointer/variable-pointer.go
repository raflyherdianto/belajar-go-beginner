package main

import "fmt"

func main() {
	var a int = 10
	var ip *int = &a

	*ip = 20

	fmt.Printf("Address of a variable: %x\n", &a)

	fmt.Printf("Address stored in ip variable: %x\n", ip)

	fmt.Printf("Value of *a variable: %d\n", a)

}
