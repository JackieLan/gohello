package main

import (
	"fmt"
	"github.com/JackieLan/gogreetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := gogreetings.Hello("joe")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := gogreetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
