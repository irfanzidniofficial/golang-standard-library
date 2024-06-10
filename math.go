package main

import (
	"fmt"
	"math"
)

func main() {
	// Membulatkan ke atas
	fmt.Println(math.Ceil(1.40))
	// Membulatkan ke bawah
	fmt.Println(math.Floor(1.60))
	
	// Membulatkan yang terdekat
	fmt.Println(math.Round(1.60))
	// Mencari nilai yang terbesar
	fmt.Println(math.Max(10,11))
	// Mencari nilai yang terkecil
	fmt.Println(math.Min(10,11))
}