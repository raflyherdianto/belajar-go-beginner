package main

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banana", "melon"}
	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	fmt.Println(fruits)
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))

	fmt.Println(aFruits)
	fmt.Println(len(aFruits))
	fmt.Println(cap(aFruits))

	fmt.Println(bFruits)
	fmt.Println(len(bFruits))
	fmt.Println(cap(bFruits))

	aFruits[0] = "pineapple"

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)

	var cFruits = append(bFruits, "papaya")

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(cFruits)

	var dFruits = append(cFruits, "pomegranate")

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(cFruits)
	fmt.Println(dFruits)

	var eFruits = append(dFruits, "pear")

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(cFruits)
	fmt.Println(dFruits)
	fmt.Println(eFruits)

	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))
}
