package main

import "fmt"

func main() {
	var data = []string{"satu", "dua", "tiga", "empat", "lima"}

	filter(data, func(each string) bool {
		return each != "dua"
	})
}

func filter(data []string, callback func(string) bool) []string {
	var result []string

	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}

	fmt.Println(result)
	return result
}
