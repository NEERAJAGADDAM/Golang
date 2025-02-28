package main

import (
	"fmt"
)


func addContact(directory map[string]string, name, phone string) {
	directory[name] = phone
	fmt.Println("Contact added successfully.")
}


func deleteContact(directory map[string]string, name string) {
	if _, exists := directory[name]; exists {
		delete(directory, name)
		fmt.Println("Contact deleted successfully.")
	} else {
		fmt.Println("Contact not found.")
	}
}

func searchContact(directory map[string]string, name string) {
	if phone, exists := directory[name]; exists {
		fmt.Printf("Contact found: %s - %s\n", name, phone)
	} else {
		fmt.Println("Contact not found.")
	}
}


func displayContacts(directory map[string]string) {
	if len(directory) == 0 {
		fmt.Println("Phone directory is empty.")
		return
	}
	fmt.Println("Phone Directory:")
	for name, phone := range directory {
		fmt.Printf("%s: %s\n", name, phone)
	}
}

func main() {
	
	phoneDirectory := make(map[string]string)

	for {
		fmt.Println("\nPhone Directory Menu:")
		fmt.Println("1. Add Contact")
		fmt.Println("2. Delete Contact")
		fmt.Println("3. Search Contact")
		fmt.Println("4. Display Contacts")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var name, phone string
			fmt.Print("Enter name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter phone number: ")
			fmt.Scanln(&phone)
			addContact(phoneDirectory, name, phone)
		case 2:
			var name string
			fmt.Print("Enter name to delete: ")
			fmt.Scanln(&name)
			deleteContact(phoneDirectory, name)
		case 3:
			var name string
			fmt.Print("Enter name to search: ")
			fmt.Scanln(&name)
			searchContact(phoneDirectory, name)
		case 4:
			displayContacts(phoneDirectory)
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
