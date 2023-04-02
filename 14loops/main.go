package main

import "fmt"

func main() {
	fmt.Println("welcome to the loops ")
	days := []string{"sunday", "monday", "tuesday"}
	fmt.Printf("type: %T \n", days)
	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	fmt.Println("--------_ANOTHER WAY OF FOR LOOOPS------------")
	for i := range days {
		fmt.Println(days[i])
	}
	fmt.Println("--------_ANOTHER WAY OF FOR LOOOPS------------")
	for i, j := range days {
		fmt.Printf("index--> %v and value--> %v \n", i, j)
	}
	fmt.Println("--------_ANOTHER WAY OF while loop in the form of for  LOOOPS------------")
	condition := 1
	for condition < 11 {
		if condition == 3 {
			condition++
			fmt.Println("wow it is 3 so continue ")
			continue
		}
		if condition == 5 {
			fmt.Println("this is 5")
			break
		}

		if condition == 8 {
			goto jump
		}
		fmt.Println(condition)
		condition++
	}

jump:
	fmt.Println("this is 8")
}
