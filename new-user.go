package main

import "fmt"

// user struct
type User struct {
	firstName string
	lastName  string
	birthDate string
}

// write data in struct
func NewUser(fName string, lName string, bDate string) *User {

	user := User{
		firstName: fName,
		lastName:  lName,
		birthDate: bDate,
	}
	return &user
}

// output data from struct (struct method)
func (user *User) outputData() {
	fmt.Printf("My name is %v %v and born on %v", user.firstName, user.lastName, user.birthDate)
}
