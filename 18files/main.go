package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to golang files")
	contents := "this need to be added in the files --Soumya"
	file, err := os.Create("./mylog.txt")

	if err != nil {
		panic(err)

	}
	length, err := io.WriteString(file, contents)

	if err != nil {
		panic(err)
	}
	fmt.Println("length is : ", length)
	defer file.Close()

	readfile("./mylog.txt")

}

func readfile(name string) {
	databyte, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	fmt.Println("files in data bytes : ", string(databyte))
}
