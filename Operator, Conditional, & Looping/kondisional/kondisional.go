package main

import "fmt"

func main() {
	var nilai int = 90

	if nilai >= 90 {
		fmt.Println("Nilai anda A")
	} else if nilai >= 80 {
		fmt.Println("Nilai anda B")
	} else if nilai >= 70 {
		fmt.Println("Nilai anda C")
	} else if nilai >= 60 {
		fmt.Println("Nilai anda D")
	} else {
		fmt.Println("Nilai anda E")
	}

	// Switch Case
	switch nilai {
	case 90:
		fmt.Println("Nilai anda A")
	case 80:
		fmt.Println("Nilai anda B")
	case 70:
		fmt.Println("Nilai anda C")
	case 60:
		fmt.Println("Nilai anda D")
	default:
		fmt.Println("Nilai anda E")
	}

	// Switch Case dengan kondisi
	switch {
	case nilai >= 90:
		fmt.Println("Nilai anda A")
	case nilai >= 80:
		fmt.Println("Nilai anda B")
	case nilai >= 70:
		fmt.Println("Nilai anda C")
	case nilai >= 60:
		fmt.Println("Nilai anda D")
	default:
		fmt.Println("Nilai anda E")
	}
}
