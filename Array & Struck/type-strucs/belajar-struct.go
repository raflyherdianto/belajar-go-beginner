package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	var user User
	user.Name = "John Wick"
	user.Age = 27

	fmt.Println(user.Name)
	fmt.Println(user.Age)

	var user2 = User{Name: "John Wick", Age: 27}
	fmt.Println(user2.Name)
	fmt.Println(user2.Age)

	var user3 = User{Name: "John Wick"}
	fmt.Println(user3.Name)
	fmt.Println(user3.Age)
}
