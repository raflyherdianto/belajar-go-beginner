package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonString = `{"Name": "John Wick", "Age": 27}`
	var jsonData = []byte(jsonString)

	var data map[string]interface{}
	json.Unmarshal(jsonData, &data)
	fmt.Println("user :", data["Name"])
	fmt.Println("age  :", data["Age"])

	var data2 interface{}
	json.Unmarshal(jsonData, &data2)
	var decodedData = data2.(map[string]interface{})
	fmt.Println("user :", decodedData["Name"])
	fmt.Println("age  :", decodedData["Age"])

}
