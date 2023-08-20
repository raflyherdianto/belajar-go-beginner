package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var jsonString = `[
		{"Name": "John Wick", "Age": 27},
		{"Name": "Ethan Hunt", "Age": 32}
	]`
	var jsonData = []byte(jsonString)

	var data []User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user 1:", data[0].FullName)
	fmt.Println("user 2:", data[1].FullName)
}
