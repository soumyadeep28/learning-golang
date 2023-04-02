package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("this is a new video of web request")
	Publicgetrequest()
	PerformjsonPostreq()
	PerformPostform()

}

func Publicgetrequest() {
	// if P is in caps of function definition then it is the public one else it is private

	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("status code:  ", response.StatusCode)
	fmt.Println("content length :  ", response.Request.ContentLength)
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

	// another way to define a response
	var responsestring strings.Builder
	content1, _ := ioutil.ReadAll(response.Body)
	bytecount, _ := responsestring.Write(content1)

	fmt.Println("bytecount: ", bytecount)
	fmt.Println("response: ", responsestring.String())

}

func PerformjsonPostreq() {
	const myurl = "http://localhost:8000/post"

	//create fake jsonpayload
	requestbody := strings.NewReader(`
	 	{
			"coursename":"python",
			   "price": 0 ,
			"plstform" : "learncodeOnline"
		}  
	   `)
	resp, err := http.Post(myurl, "application/json", requestbody)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("output: ", string(content))
}

func PerformPostform() {
	const myurl = "http://localhost:8000/post"
	data := url.Values{}
	data.Add("name", "soumya")
	data.Add("roll", "seven")
	data.Add("sec", "A")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
