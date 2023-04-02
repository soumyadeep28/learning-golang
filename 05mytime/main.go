package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time of golang")
	timenow := time.Now()
	fmt.Println("curent time is now")
	fmt.Println(timenow)

	fmt.Println("added the FORMAT option")
	fmt.Println(timenow.Format("01-02-2006 Monday 15:04:05"))

	fmt.Println("added the Date option")
	createddate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createddate)

}
