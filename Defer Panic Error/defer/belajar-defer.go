package main

import "fmt"

func main() {
	defer fmt.Println("Ini akan dijalankan terakhir")
	fmt.Println("Ini akan dijalankan pertama")
}
