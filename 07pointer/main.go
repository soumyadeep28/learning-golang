package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class of pointer ")
	// var ptr *int
	// fmt.Println("the value of pointer : ", ptr)

	var1 := 10
	var ptr = &var1 //this is pointer referenced
	fmt.Println("address of the pointer : ", ptr)
	fmt.Println("value of the pointer : ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("new value of the variable : ", var1)

}
