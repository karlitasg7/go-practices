package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func saveContactToFile(contacts []Contact) error {

	file, err := os.Create("contacts.json")

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(contacts)

	if err != nil {
		return err
	}

	return nil

}

func loadContactsFromFile(contacts *[]Contact) error {

	file, err := os.Open("contacts.json")

	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&contacts)

	if err != nil {
		return err
	}

	return nil

}

func main() {

	var contacts []Contact
	err := loadContactsFromFile(&contacts)

	if err != nil {
		fmt.Println("Error on loading contacts", err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("1. Add Contact \n",
			"2. List Contacts \n",
			"3. Exit \n",
			"Enter your choice: ")

		var option int
		_, err := fmt.Scanln(&option)

		if err != nil {
			fmt.Println("Invalid option", err)
			return
		}

		switch option {

		case 1:
			var c Contact
			fmt.Print("Enter name: ")
			c.Name, _ = reader.ReadString('\n')

			fmt.Print("Enter email: ")
			c.Email, _ = reader.ReadString('\n')

			fmt.Print("Enter phone: ")
			c.Phone, _ = reader.ReadString('\n')

			contacts = append(contacts, c)

			err := saveContactToFile(contacts)

			if err != nil {
				fmt.Println("Error on saving contacts", err)
			}

		case 2:

			fmt.Println("Contacts:")
			fmt.Println("=====================================")

			for index, c := range contacts {
				fmt.Printf("%d. Name: %s Email: %s Phone: %s \n", index+1, c.Name, c.Email, c.Phone)
			}
			fmt.Println("=====================================")

		case 3:
			return

		default:
			fmt.Println("Invalid option")

		}
	}

}
