package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Ini akan dijalankan pertama")

	os.Exit(1)

	fmt.Println("Ini tidak akan pernah dijalankan")
}
