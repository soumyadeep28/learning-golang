package main

import "fmt"

func main() {
	fmt.Println("this is th maps in golang ")
	var languages = make(map[string]string)

	languages["JS"] = "javascript"
	languages["PY"] = "Python"
	languages["JAVA"] = "JAVA"
	fmt.Println("list of all the languages: ", languages)

	fmt.Println(" languages js: ", languages["JS"])

	// delete the language
	delete(languages, "JAVA")
	fmt.Println("list of all the languages: ", languages)

	//loops in maps
	for key, value := range languages {
		fmt.Printf("\n this is the key %v and this is the key: %v", key, value)
		fmt.Println(key, "-->", value)
	}

}
