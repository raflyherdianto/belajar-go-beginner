package main

import "fmt"

func main() {
	// Perulangan dengan for
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Perulangan dengan while
	var i = 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// Perulangan dengan for tanpa argumen
	var j = 0
	for {
		fmt.Println(j)
		j++
		if j == 5 {
			break
		}
	}

	// Perulangan dengan for range
	var k = []string{"a", "b", "c", "d", "e"}
	for i, v := range k {
		fmt.Println("index", i, "=", v)
	}

	// Perulangan dengan for range tanpa index
	var l = []string{"a", "b", "c", "d", "e"}
	for _, v := range l {
		fmt.Println(v)
	}

	// Perulangan dengan for range pada map
	var m = make(map[string]string)
	m["nama"] = "Rizky"
	m["alamat"] = "Bekasi"
	for k, v := range m {
		fmt.Println(k, "=", v)
	}
}
