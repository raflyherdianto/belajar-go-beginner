package main

import "fmt"

func main() {
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"])
	fmt.Println("mei", chicken["mei"])

	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	fmt.Println("januari", chicken2["januari"])
	fmt.Println("mei", chicken2["mei"])

	for key, val := range chicken2 {
		fmt.Println(key, "  \t:", val)
	}

	delete(chicken2, "januari")
	fmt.Println(len(chicken2))
	fmt.Println(chicken2)
}
