// hello
package main

const spanish = "Spanish"
const french = "French"
const spanishHelloPrefix = "Hella, "
const frenchHelloPrefix = "Bonjour, "
const englishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix = englishHelloPrefix
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}

	return prefix + name
}
