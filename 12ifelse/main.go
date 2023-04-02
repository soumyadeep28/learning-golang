package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("this is a lession of ifelse case ")

	var result string
	logincount := bufio.NewReader(os.Stdin)
	input, _ := logincount.ReadString('\n')
	numrating, _ := strconv.ParseInt(strings.TrimSpace(input), 64, 10)
	if numrating < 10 {
		result = "this is not a regular user"
	} else if numrating > 10 && numrating < 20 {
		result = "medium user"
	} else {
		result = "this is a reguler one "
	}

	fmt.Println(result)
}
