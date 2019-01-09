package main

import "fmt"

const helloPrefix = "Hello "
const japanese = "Japanese"
const japaneseHelloPrefix = "konnichiwa "
const german = "German"
const germanHelloPrefix = "Hallo "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := helloPrefix

	switch language {
	case japanese:
		prefix = japaneseHelloPrefix
	case german:
		prefix = germanHelloPrefix
	}

	return prefix + name
}
func main() {
	fmt.Println(Hello("Sarah", "English"))
}
