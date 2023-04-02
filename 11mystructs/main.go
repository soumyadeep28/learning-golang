package main

import "fmt"

func main() {
	fmt.Println("this is the new era of Golang ")
	// no inheritance in golang : no super , super parent

	// creating kind of a class
	hitesh := User{"Hitesh", "hitech@go.dev", true, 30}
	fmt.Println(hitesh)
	fmt.Printf("thers are the details: %+v ", hitesh)

}

type User struct {
	// first one should be in the capital letter
	Name   string
	Email  string
	Status bool
	Age    int
}
