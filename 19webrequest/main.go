package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("welcome to webrequest lessions ")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response of type : %T", response)
	fmt.Println(response)

	defer response.Body.Close() //callers responsibility to close that
	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(resp)
	fmt.Println(content)

}
