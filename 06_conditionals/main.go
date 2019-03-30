package main

import "fmt"

func main() {
	color := "green"

	if color == "blue" {
		fmt.Println("The color is blue.")
	} else if color == "red" {
		fmt.Println("The color is red.")
	} else {
		fmt.Println("I don't know which color it is.")
	}

	switch color {
	case "blue":
		fmt.Println("The color is blue.")
	case "red":
		fmt.Println("The color is red.")
	default:
		fmt.Println("I chose purple for you.")
	}

}
