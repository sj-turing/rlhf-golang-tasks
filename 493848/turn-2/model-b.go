package main

import "fmt"

// PrintMessage is a callback function that takes a message and prints it.
type PrintMessage func(string)

// PrintNumber is a callback function that takes an integer and prints it.
type PrintNumber func(int)

// ComposeMessageAndNumber composes PrintMessage and PrintNumber functions.
func ComposeMessageAndNumber(messagePrinter PrintMessage, numberPrinter PrintNumber) func(string, int) {
	return func(message string, number int) {
		messagePrinter(message)
		numberPrinter(number)
	}
}

// PrintMessageFunc is an implementation of PrintMessage that prints a message.
func PrintMessageFunc(message string) {
	fmt.Println(message)
}

// PrintNumberFunc is an implementation of PrintNumber that prints a number.
func PrintNumberFunc(number int) {
	fmt.Println(number)
}

func main() {
	// Compose the PrintMessageFunc and PrintNumberFunc
	composedFunction := ComposeMessageAndNumber(PrintMessageFunc, PrintNumberFunc)

	// Call the composed function with a message and a number
	composedFunction("Hello, World!", 42)
	// Output:
	// Hello, World!
	// 42
}
