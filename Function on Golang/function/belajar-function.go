package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func sayHelloTo(name string) {
	fmt.Println("Hello", name)
}

func sayHelloReturn(name string) string {
	return "Hello " + name
}

func main() {
	// Function
	sayHello()
	sayHelloTo("Eko")
	result := sayHelloReturn("Budi")
	fmt.Println(result)
}
