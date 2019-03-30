package main

import (
	"fmt"
	"math"

	"github.com/juliomsilva/go_crash_course/03_packages/strutil"
)

func main() {
	fmt.Printf("This is 3ˆ2: %v\n ", math.Pow(3, 2))
	fmt.Println(strutil.Reverse("Rerver this!"))
}
