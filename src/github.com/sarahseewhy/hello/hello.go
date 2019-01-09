package main

import (
	"fmt"
)

const englishPrefix = "Hello "
const japanese = "Japanese"
const japaneseHelloPrefix = "konnichiwa "
const german = "German"
const germanHelloPrefix = "Hallo "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case japanese:
		prefix = japaneseHelloPrefix
	case german:
		prefix = germanHelloPrefix
	default:
		prefix = englishPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Sarah", "English"))
}
