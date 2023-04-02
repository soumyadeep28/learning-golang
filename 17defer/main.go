package main

import "fmt"

/*
this defer key word implies that th line where defer is used it will goto the end of the execution in a stack order
*/
func main() {
	fmt.Println("implementations of defer: ")
	fmt.Println("run go")
	defer fmt.Println("lang")

	fmt.Println("this is new line")

	defer fmt.Println("lang2")
	fmt.Println("this is new line")

	printMyfunc()

}

func printMyfunc() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
