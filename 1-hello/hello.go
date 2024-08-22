package hello

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	german             = "German"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	germanHelloPrefix  = "Hallo, "
)

// Hello greets according to provided name and language
func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	return GreetingPrefix(language) + name + "!"
}

// GreetingPrefix returns the correct greeting depending on the language
func GreetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case german:
		prefix = germanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", "English"))
}
