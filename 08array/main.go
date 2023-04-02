package main

import "fmt"

func main() {
	fmt.Println("Welcome to the array of golang")
	var frouts [4]string
	frouts[0] = "mango1"
	frouts[1] = "mango2"
	frouts[2] = "mango3"
	frouts[3] = "mango4"

	fmt.Println("all the frouts: ", frouts)

	var bag [4]string
	bag[0] = "mango1"
	// this is blank and we havent passed the value
	bag[2] = "mango3"
	bag[3] = "mango4"

	fmt.Println("all the frouts: ", bag)

	// you can see the differece
	// all the frouts:  [mango1 mango2 mango3 mango4]
	// all the frouts:  [mango1  mango3 mango4]  <-- here space is double after the mango1 and mango2

	fmt.Println("there is the differnce b/w array in goalng with other language")
	fmt.Println("the size of the array frouts and bags are  : ", len(frouts), len(bag))
	fmt.Println("you can see in the both cases length is 4 . ")

	fmt.Println("Another way to define the array ")
	var myfilelist = [20]string{"file1", "file2"}
	fmt.Println("the fiellist: ", myfilelist)
	fmt.Println("length of the fiellist: ", len(myfilelist))

}
