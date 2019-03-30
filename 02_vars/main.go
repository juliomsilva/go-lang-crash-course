package main

import "fmt"

func main() {
	/*
	 * ASSIGNEMENT
	 *
	 * var, const
	 * Operators: default =, shorthand :=
	 *
	 * SCOPE RULES
	 *
	 * No short hand out of functions bodies.
	 *
	 * MAIN TYPES
	 *
	 * string
	 * bool
	 *
	 * int, int8, int16, int32, int64
	 * byte: alias for uint8
	 *
	 * uint, uint8, uint16, uint32, uint64, uintptr
	 * rune: alias for int32
	 *
	 * float32, float64
	 * comples64, comples128
	 */

	// name := "Júlio"
	// name = "Júlio Silva"
	// email := "toomuch.js@gmail.com"

	name, email := "Júlio Silva", "toomuch.js@gmail.com"

	var age = 31
	var isCool = true
	var weigth float32 = 79.2
	height := 1.81

	fmt.Println(name, age, email, isCool, weigth, height)
	fmt.Printf("%T, %T, %T, %T, %T, %T\n", name, age, email, isCool, weigth, height)

}
