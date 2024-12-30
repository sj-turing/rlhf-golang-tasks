package main

import (
	"fmt"
	"strings"
)

type Callback interface {
	Invoke(data interface{}) error
}

type GreetCallback struct {
	Msg string
}

func (gc GreetCallback) Invoke(data interface{}) error {
	name, ok := data.(string)
	if !ok {
		return fmt.Errorf("Invalid data type for GreetCallback")
	}
	fmt.Printf("%s, %s!\n", gc.Msg, name)
	return nil
}

func Compose(callbacks ...Callback) Callback {
	return composeCallback{callbacks}
}

type composeCallback struct {
	callbacks []Callback
}

func (cc composeCallback) Invoke(data interface{}) error {
	for _, cb := range cc.callbacks {
		if err := cb.Invoke(data); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// Define some callbacks
	greetCallback := GreetCallback{"Hello"}
	upperCaseCallback := func(data interface{}) error {
		str, ok := data.(string)
		if !ok {
			return fmt.Errorf("Invalid data type for upperCaseCallback")
		}
		data = strings.ToUpper(str)
		return nil
	}

	// Compose callbacks
	composedCallback := Compose(greetCallback, upperCaseCallback)

	// Invoke the composed callback
	err := composedCallback.Invoke("world")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
