package main

import "fmt"

func main() {
	fmt.Println("this is a method package : ")

	// creating kind of a class
	hitesh := User{"Hitesh", "hitech@go.dev", true, 30, 20}
	fmt.Println(hitesh)
	fmt.Printf("thers are the details: %+v \n ", hitesh)
	hitesh.Getstatus()
	hitesh.getemail()
	fmt.Println(hitesh)
	fmt.Println("this is to show when we pass the address of the object it is changing the value int he object")
	hitesh.getgmail()
	fmt.Println(hitesh)
}

type User struct {
	// first one should be in the capital letter
	Name   string
	Email  string
	Status bool
	Age    int
	oneAge int
}

func (u User) Getstatus() {
	fmt.Println("Status of the uSer : ", u.Status)

}
func (u User) getemail() {
	u.Email = "abhi@gmail.com"
	fmt.Println(u.Email)
}
func (u *User) getgmail() {
	u.Email = "abhi@gmail.com"
	fmt.Println(u.Email)
}
