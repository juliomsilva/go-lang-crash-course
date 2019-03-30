package main

import "fmt"

func main() {
	// Long method.
	i := 1
	for i <= 10 {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Printf("\n")

	// Short method.
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")

	// FizzBuzz
	for i := 1; i <= 1000; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}
