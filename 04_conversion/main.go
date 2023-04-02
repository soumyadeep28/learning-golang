package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("'Welcome to tthe world fo programming : ")
	fmt.Println("'enter the rating for ouyr course : ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("thanks for rating ", input)
	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		numrating = numrating + 1
		fmt.Println("added 1 with rating ")
		fmt.Println("now rating ", numrating)
	}

}
