package main

import (
	"fmt"
	"greetings/greetings"
	"log"

	"rsc.io/quote"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(1)

	fmt.Println("Hello, World!")
	fmt.Println(quote.Glass())

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}

//https://go.dev/doc/tutorial/call-module-code
