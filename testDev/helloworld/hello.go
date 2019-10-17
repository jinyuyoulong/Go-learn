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

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	prefix := englishHelloPrefix
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
