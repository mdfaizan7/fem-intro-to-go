package main

import "fmt"

// User is a user type
type User struct {
	ID                         int
	firstName, lastName, email string
}

// Group represents a group of user
type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

func describeUser(u User) string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.firstName, u.lastName, u.email)
	return desc
}

func describeGroup(g Group) string {
	if len(g.users) >= 2 {
		g.spaceAvailable = false
	}
	desc := fmt.Sprintf("This user group has %d users. The newest user is %s %s. Accepting New Users: %t", len(g.users), g.newestUser.firstName, g.newestUser.lastName, g.spaceAvailable)
	return desc
}

func main() {
	u := User{
		ID:        1,
		firstName: "Tony",
		lastName:  "Stark",
		email:     "tonystark@ironman.com"}
	u2 := User{
		ID:        2,
		firstName: "Peter",
		lastName:  "Parker",
		email:     "peterparker@spiderman.com"}
	u3 := User{
		ID:        3,
		firstName: "Natasha",
		lastName:  "Romanoff",
		email:     "natasharomanoff@blackwidow.com"}

	g := Group{
		role:           "admin",
		users:          []User{u, u2, u3},
		newestUser:     u3,
		spaceAvailable: true}

	fmt.Println(describeUser(u))
	fmt.Println(describeGroup(g))
	fmt.Println(g)
}
