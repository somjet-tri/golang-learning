package main

import "fmt"

func main() {
	// create a slice of ints from 0 through 10
	numbers := []int{}
	for i := 0; i <= 10; i++ {
		numbers = append(numbers, i)
	}
	fmt.Println(numbers)

	for _, n := range numbers {
		if n%2 == 0 {
			fmt.Printf("%v is even.\n", n)
		} else {
			fmt.Printf("%v is odd.\n", n)
		}
	}
}
