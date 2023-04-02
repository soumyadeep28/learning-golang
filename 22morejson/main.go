package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //this is to create a alias
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags ,omitempty"`
}

func main() {
	fmt.Println("welcome to video of json : ")
	EncodingJson()
}

func EncodingJson() {
	lcocourses := []course{
		{"botcamp1", 200, "mywebsite.com", "abc123", []string{"boot", "camp"}},
		{"botcamp2", 400, "mywebsite.com", "abc123", []string{"boot1", "camp"}},
		{"botcamp3", 400, "mywebsite.com", "abc123", []string{"boot2", "camp"}},
		{"botcamp4", 1200, "mywebsite.com", "abc123", nil},
	}
	//finaljson, err := json.Marshal(lcocourses)
	//this marshalindent is to design in a mode of labeling into structured way
	finaljson, err := json.MarshalIndent(lcocourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", finaljson)
}
