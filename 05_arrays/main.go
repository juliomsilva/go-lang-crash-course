package main

import "fmt"

func main() {
	var fruitArr [2]string
	fruitArr[0], fruitArr[1] = "Apple", "Orange"

	vegetableArr := []string{"Lettuce", "Onion", "Cabbage"}

	fmt.Println(fruitArr)
	fmt.Println(vegetableArr)

	fmt.Println(vegetableArr[1:2])
}
