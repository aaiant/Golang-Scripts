package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(firstName, lastName, birthDate, email, password string) (*Admin, error) {
	if firstName == "" || lastName == "" || birthDate == "" || email == "" || password == "" {
		return nil, errors.New("first name, last name, birthdate, email and password are required")
	}

	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: firstName,
			lastName:  lastName,
			birthDate: birthDate,
			createdAt: time.Now(),
		},
	}, nil
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func (user User) OutputUserDetails() {
	fmt.Println(user.firstName, user.lastName, user.birthDate)
}

func (user *User) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}
