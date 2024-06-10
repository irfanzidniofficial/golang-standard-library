package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("true")
	if(err != nil){
		fmt.Println("Error", err.Error())
	} else{
		fmt.Println(result)
	}

	resultInt, err := strconv.Atoi("10000")
	if err != nil{
		fmt.Println("Error", err.Error())

	} else {
		fmt.Println(resultInt)
	}

	binary:= strconv.FormatInt(999,2)
	fmt.Println(binary)

	
		
}

