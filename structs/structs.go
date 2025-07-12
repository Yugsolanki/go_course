package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Enter first name: ")
	lastName := getUserData("Enter last name: ")
	birthDate := getUserData("Enter birth date: ")

	appUser, err := user.New(firstName, lastName, birthDate)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	// var appUser user.User = user.User{
	// 	firstName: firstName,
	// 	lastName:  lastName,
	// 	birthDate: birthDate,
	// 	createdAt: time.Now(),
	// }
	// var appUser user = user{}	empty struct
	// var appuser user = user{firstName: firstName}	rest will values will be assiged default value

	adminUser := user.NewAdmin("yug@gmail.com", "123456")
	adminUser.User.OutputUserDetails() // even this work
	adminUser.User.ClearUserName()
	adminUser.OutputUserDetails() // even this work with the .User

	fmt.Println("------------")

	// fmt.Println(appUser)
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

// func outputUserDetails(u *user) {
// fmt.Println((*u).firstName)	with de-referencing, but no need to do it as golang allows for a shortcut as below
// 	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
// }

func getUserData(text string) string {
	fmt.Print(text)
	var value string
	fmt.Scanln(&value)
	return value
}
