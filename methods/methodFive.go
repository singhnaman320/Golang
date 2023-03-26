package main

import "fmt"

type User struct {
	name   string
	email  string
	status bool
	age    int
}

func (u User) GetStatus() {

	fmt.Println("\nStatus of user:", u.status)
}
func (u User) newMail() {

	u.email = "singhnaman320@gmail.com"

	fmt.Println("New email is:", u.email)
}

func main() {

	fmt.Println("Struct in Golang")

	//There is no inheritance no super no parent in Golang

	thisUser := User{"naman", "naman@123", true, 67}

	fmt.Println(thisUser)

	fmt.Printf("Given user details are: %v\n", thisUser)

	fmt.Printf("Name is %v and email is %v", thisUser.name, thisUser.email)

	thisUser.GetStatus()

	thisUser.newMail()

	fmt.Printf("Name is %v and email is %v", thisUser.name, thisUser.email)
}
