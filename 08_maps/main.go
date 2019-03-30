package main

import "fmt"

func main() {
	// Define map.
	emails := make(map[string]string)
	skills := map[string]string{"bob": "numbers", "sharon": "peoples"}

	// Assign key values.
	emails["bob"] = "bob@email.com"
	emails["sharon"] = "sharon@email.com"

	// Bob got fired.
	delete(emails, "bob")
	delete(skills, "bob")

	fmt.Println(emails)
	fmt.Println(skills)
}
