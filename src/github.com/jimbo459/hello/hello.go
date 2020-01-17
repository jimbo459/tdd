package main

import "fmt"

const defaultPrefix = "Hello, "
const frenchPrefix = "Bonjour, "
const spanishPrefix = "Hola, "

const french = "French"
const spanish = "Spanish"

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = defaultPrefix
	}
	return
}

func hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name

}

func main() {
	fmt.Println(hello("name", "language"))
}
