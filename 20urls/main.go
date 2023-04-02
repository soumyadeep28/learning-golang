package main

import (
	"fmt"
	"net/url"
)

const urls = "http://lco.dev:3000/learn?coursename=reactjs&paymentid=jdfonsonosnfonsd"

func main() {
	fmt.Println("welcome to urls")
	fmt.Println(urls)

	//parsing
	result, _ := url.Parse(urls)
	fmt.Println(result.Scheme)
	fmt.Println(result.Port())
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	//param in better format
	qparam := result.Query()
	fmt.Println(qparam)
	fmt.Printf("type %T", qparam)

	//now construct an url :
	partsofurl := &url.URL{
		Scheme:  "https",
		Host:    "learncodeonline.com",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	rawurl := partsofurl.String()
	fmt.Println(rawurl)

}
