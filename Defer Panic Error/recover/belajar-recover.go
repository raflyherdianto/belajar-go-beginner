package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	defer catch()
	var input string
	fmt.Print("Masukkan nama anda: ")
	fmt.Scanln(&input)

	err := validate(input)
	if err != nil {
		panic(err.Error())
		fmt.Println(err.Error())
	}

	fmt.Println("halo", input)
}

func validate(input string) error {
	if strings.TrimSpace(input) == "" {
		return errors.New("nama tidak boleh kosong")
	}
	return nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}
