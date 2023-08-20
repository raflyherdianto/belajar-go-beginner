package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "adalah angka")
	} else {
		fmt.Println(input, "bukan angka")
		fmt.Println(err.Error())
	}
}
