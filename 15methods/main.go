package main

import "fmt"

// Here copied this code from Struct file (10mystructs)
func main() {
	// Struct in Go is a user-defined composite data type that allows you to group multiple variables (fields) of different data types under a single name.

	// Struct is alternative version of classes, and we don't have classes in go lang. No inheritance in go lang. No super or no Parent in go lang.
	fmt.Println("Stuct in go lang")

	sad := User{"Saddam", "sad@gmail.com", true, 28}
	fmt.Println(sad)                                                 // {Saddam sad@gmail.com true 28}
	fmt.Printf("Sad details are: %+v\n", sad)                        // Sad details are: {Name:Saddam Email:sad@gmail.com Status:true Age:28}
	fmt.Printf("Name is %v and email is %v.\n", sad.Name, sad.Email) // Name is Saddam and email is sad@gmail.com.

	// Here GetStatus method call
	sad.GetStatus()

	// Here call NewEmail method
	sad.NewEmail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

// Here now update the existing email
func (u User) NewEmail() {
	u.Email = "test@gmail.com"
	fmt.Println("Now email of this user is:", u.Email)
}
