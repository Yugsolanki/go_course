package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string // change it to `FirstName` to also export the field
	lastName  string
	birthDate string
	createdAt time.Time
}

// Type of inheritence
type Admin struct {
	email    string
	password string
	User
	// baseUser User 	:- you also assign name
	// BaseUser User	:- to also export
}

// func (user) outputUserDetails(u *user) {41``
// 	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
// }
// for the above one you will need to pass an argument like appUser.outputUserDetails(&appuser)

// no need to pass an argument the `u` is called receiver, but can also pass additional arugments apart from `u`
// `u` user is passed a copy use `u *user` to pass reference
func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

// Mutation function: func that mute data of struct rather than just outputing them
// user reference is used here
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// Constructor/Creation function
func New(firstName, lastName, birthDate string) (*User, error) {

	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("Got empty string")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "Admin",
			birthDate: "3/5/2003",
			createdAt: time.Now(),
		},
	}
}
