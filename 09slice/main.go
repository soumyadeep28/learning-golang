package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Enter the world of go programming  : slices ")
	// we are using the basic initialization of slices

	var bucket = []string{"apple", "orange", "goava"}
	fmt.Println("this is the bucket : ", bucket)
	fmt.Printf("this is the type of the bucket : %T \n", bucket)

	// this is the method to add the param

	bucket = append(bucket, "banana", "coconut")
	fmt.Println("this is the bucket : ", bucket)
	fmt.Printf("this is the type of the bucket : %T \n", bucket)
	// this is to slicing the variable

	bucket = append(bucket[1:]) //this will start form 1 index and then go along
	fmt.Println("this is the sliced the version of the list : ", bucket)

	bucket = append(bucket[1:3]) //this will start form 1 index and then go along
	fmt.Println("this is the sliced the version of the list : ", bucket)

	// now we are using make() and new() function to create the new params

	highscores := make([]int, 4) // this will allocate the memory of 4 of intger

	highscores[0] = 10
	highscores[1] = 11
	highscores[2] = 12
	highscores[3] = 13
	fmt.Println(highscores)
	highscores = append(highscores, 14, 15, 16)
	fmt.Println(highscores)
	sort.Ints(highscores)
	fmt.Println(highscores)

	// this returns the default value of the true or false
	fmt.Println(highscores)
	var1 := sort.IntsAreSorted(highscores)
	fmt.Println(var1)

	//How to remove a value from slices
	var courses = []string{"react", "node", "js", "python", "java", "python", "c++"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
