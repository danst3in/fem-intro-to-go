package main

import "fmt"

// A User is a user type
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

var u = User{
	ID:        1,
	FirstName: "Marylin",
	LastName:  "Monroe",
	Email:     "marilynmonroe@gmail.com",
}

func updateEmail(u *User) {
	u.Email = "anotheremail@gmail.com"
	fmt.Println("inside update: %s", u.Email)
}

func main() {
	fmt.Println("OG Email: %s", u.Email)

	updateEmail(&u)
	fmt.Println("user email after update: %s", u.Email)
	fmt.Println("user after update: %s", u)
}
