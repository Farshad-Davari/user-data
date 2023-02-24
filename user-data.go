package main

func main() {

	var user *User

	firstName := userInputs("Please enter your first name: ")
	lastName := userInputs("Please enter your last name: ")
	birthDate := userInputs("Please enter your birth date(DD/MM/YYYY): ")

	user = NewUser(firstName, lastName, birthDate)

	user.outputData()
}
