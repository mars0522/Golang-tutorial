package main

import (
	"fmt"
	"time"
)

type userInfo struct {
	firstName    string
	userLastName string
	dob          string
	createdAt    time.Time
}

// This is go method
func (u *userInfo) PrintUserInfo() {
	fmt.Println("firstName: ", u.firstName)
	fmt.Println("lastName: ", u.userLastName)
	fmt.Println("DOB: ", u.dob)
}

func main() {
	userFirstName := getUserData("Enter user first name: ")
	userLastName := getUserData("Enter user last name: ")
	DOB := getUserData("Enter user data of birth: ")

	user := userInfo{
		firstName:    userFirstName,
		userLastName: userLastName,
		dob:          DOB,
		createdAt:    time.Now(),
	}

	PrintUserInfo(&user)
	user.PrintUserInfo()

}

func PrintUserInfo(user *userInfo) {
	fmt.Printf("User Info:\n")
	fmt.Printf("First Name: %s\n", user.firstName)
	fmt.Printf("Last Name: %s\n", user.userLastName)
	fmt.Printf("Date of Birth: %s\n", user.dob)
	fmt.Printf("Created At: %s\n", user.createdAt)
	fmt.Printf("-----------------------------\n")
}
func getUserData(promtText string) string {
	fmt.Print(promtText)
	var value string
	fmt.Scan(&value)
	return value
}
