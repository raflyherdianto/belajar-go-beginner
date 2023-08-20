package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Masukkan nama anda: ")
	fmt.Scanln(&input)

	err := validate(input)
	if err != nil {
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
