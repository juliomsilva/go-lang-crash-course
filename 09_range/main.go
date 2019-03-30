package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 32, 11, 2}

	// Using index.
	for i, id := range ids {
		fmt.Printf("%d] ID: %d\n", i, id)
	}

	// Not using index.
	for _, id := range ids {
		fmt.Printf("ID: %d. ", id)
	}
	fmt.Println()

	// Range with Map.
	skills := map[string]string{"bob": "numbers", "sharon": "peoples"}

	for k, v := range skills {
		fmt.Printf("%s is skilled with %s.\n", k, v)
	}

	for _, v := range skills {
		fmt.Printf("We are skilled with %s.\n", v)
	}
}
