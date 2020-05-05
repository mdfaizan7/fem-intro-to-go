package main

import "fmt"

// User is a user struct type
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

func updateEmail(u *User, newEmail string) {
	u.Email = newEmail
}

func main() {
	u := User{
		ID:        1,
		FirstName: "Tony",
		LastName:  "Stark",
		Email:     "ironman@gmail.com"}

	fmt.Println(u)
	updateEmail(&u, "tonystark@ironman.com")
	fmt.Println(u)
}
