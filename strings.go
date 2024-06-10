package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Irfan Zidni", "Irfan")) // Apakah ada kata Irfan di kalimat Irfan Zidni
	fmt.Println(strings.Split("Irfan Zidni", " "))
	fmt.Println(strings.ToLower("Irfan Zidni"))
	fmt.Println(strings.ToUpper("Irfan Zidni"))
	fmt.Println(strings.Trim(" Irfan Zidni ", " "))
	fmt.Println(strings.ReplaceAll("Muhammad Irfan Zidni", "Irfan", "Van"))
}