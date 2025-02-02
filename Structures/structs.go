package main

import (
	"fmt"

	"example.com/Structures/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	appUser, error := user.New(firstName, lastName, birthDate)

	if error != nil {
		fmt.Println(error)
		return
	}

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	firstNameAdmin := getUserData("Please enter your first name: ")
	lastNameAdmin := getUserData("Please enter your last name: ")
	birthDateAdmin := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	emailAdmin := getUserData("Please enter your email: ")
	passwdAdmin := getUserData("Please enter your password: ")

	var admin *user.Admin
	admin, error = user.NewAdmin(firstNameAdmin, lastNameAdmin, birthDateAdmin, emailAdmin, passwdAdmin)

	if error != nil {
		fmt.Println(error)
		return
	}

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
