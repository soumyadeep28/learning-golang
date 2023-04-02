package main

import "fmt"

func main() {
	fmt.Println("this is a excersize of function: ")
	check_even_odd(40)
	sum1 := addition(10, 40)
	fmt.Println("this is the sum : ", sum1)

	sum1 = proadd(10, 40, 50)
	fmt.Println("this is the sum : ", sum1)
	var student = make(map[string]int)
	name, total := keypair(20, "soumya", 90, 80, 10)
	fmt.Println(name, "-->", total)
	student[name] = total
	fmt.Println(student)

}

func check_even_odd(num1 int) {
	if num1%2 == 1 {
		fmt.Println("this is odd")
	} else {
		fmt.Println("this is even  ")
	}
}
func addition(var1 int, var2 int) int {
	// here var1 and var2 are the type of the param then int is the return type of that
	return var1 + var2
}

func proadd(values ...int) int {
	sum := 0
	for i := range values {
		sum = sum + values[i]
	}

	return sum
}

// suppose you want to return 2 response one is the string and one is int

func keypair(roll int, name string, marks ...int) (string, int) {
	sum := 0
	for _, value := range marks {
		sum = sum + value
	}
	return name, sum
}
