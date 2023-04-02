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
	DecodeJson()
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

func DecodeJson() {
	jsondata := []byte(`
	{
		"coursename": "botcamp2",
		"Price": 400,
		"website": "mywebsite.com",
		"tags ": ["boot1","camp"]
}   
	`)

	var lcocourse course

	checkvalid := json.Valid(jsondata)
	if checkvalid {
		fmt.Println("json is valid ")
		json.Unmarshal(jsondata, &lcocourse)
		fmt.Printf("%#v", lcocourse)
		// here you cant print it as fmt.Println(lcocourse["Price"])
	} else {
		fmt.Println("this is not valid json file")
	}

	//another way to unmershal
	var onlinedata map[string]interface{}
	json.Unmarshal(jsondata, &onlinedata)
	fmt.Printf("\n %#v", onlinedata)
	fmt.Println(onlinedata["Price"])

	for k, v := range onlinedata {
		fmt.Printf("\n the key : %v and value of %v  if of type %T ", k, v, v)
	}

}
