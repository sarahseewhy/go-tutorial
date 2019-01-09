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

	if language == japanese {
		return japaneseHelloPrefix + name
	}

	if language == german {
		return germanHelloPrefix + name
	}

	return helloPrefix + name
}
func main() {
	fmt.Println(Hello("Sarah", "English"))
}
