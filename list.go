package main

import (
	"container/list"
	"fmt"
)

// Struktur data linked list 
func main() {
	var data *list.List = list.New()

	data.PushBack("Irfan")
	data.PushBack("Zidni")
	data.PushBack("Muhammad")

	var head *list.Element = data.Front()

	fmt.Println(head.Value) // irfan

	
	 next :=head.Next()
	fmt.Println(next.Value) // zidni

	 next =next.Next()
	fmt.Println(next.Value) // Muhammad

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) //
	}
}
