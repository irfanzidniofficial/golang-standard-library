package main

import "fmt"

func main() {
	firstName := "Irfan"
	lastName := "Zidni"

	fmt.Println("Hello '", firstName, lastName, "'")

	fmt.Printf("Hello '%s %s'\n", firstName, lastName)

}