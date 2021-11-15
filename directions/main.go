package main

import "fmt"

func main() {
	numbers := []int{}

	for index := 1; index <= 10; index++ {
		numbers = append(numbers, index)
	}

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Printf("%v is even \n", number)
		} else {
			fmt.Printf("%v is odd \n", number)
		}
	}
}
