package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	// Set pproperties of the predefines Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number...
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("Gladys")
	// If an error was retuned, print it to the console and
	// exit the program
	if err != nil {
		log.Fatal(err)
	}

	// If no erro was returned, print the returned message
	// to the console
	fmt.Println(message)
}
