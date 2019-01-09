package main

import "fmt"

const helloPrefix = "Hello "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == "Japanese" {
		return "konnichiwa " + name
	}
	return helloPrefix + name
}
func main() {
	fmt.Println(Hello("Sarah", "English"))
}
