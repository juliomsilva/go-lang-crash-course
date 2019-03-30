package main

import "fmt"

func main() {
	v := 5
	p := &v

	fmt.Println(v, p)
	fmt.Println(*p)
	fmt.Println(*&v)

	*p = 10
	fmt.Println(v)
}
