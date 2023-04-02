package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to go programming"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("rate your programming ")

	// comma ok || err ok  where it is kind of the try catch error this stores
	//the input as the output if error rthen _ .. you can store as error but
	// if error is passed then it will be marked as error

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating ", input)
	fmt.Printf("input of the variable %T ", input)

}
