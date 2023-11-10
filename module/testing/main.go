package main

import (
	"fmt"
	"log"

	"github.com/karlitasg7/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Karla")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	randomMessage := greetings.RandomGreeting("Karla")

	fmt.Println(randomMessage)

}
