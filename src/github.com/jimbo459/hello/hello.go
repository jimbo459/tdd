package main

import "fmt"

const defaultPrefix = "Hello, "
const frenchPrefix = "Bonjour, "
const spanishPrefix = "Hola, "

const french = "French"
const spanish = "Spanish"

func hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := defaultPrefix

	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(hello("name", "language"))
}
