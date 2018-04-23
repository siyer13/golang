package main

import "os"
import "fmt"
import "strconv"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Not a number")
	}

	fmt.Println(factorial(uint(num)))
}
