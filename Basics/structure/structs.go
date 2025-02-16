package main

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

type userInfo struct {
	firstName    string
	userLastName string
	dob          string
	createdAt    time.Time
}

// This is a go Method
func (u *userInfo) PrintUserInfo() {
	fmt.Println("firstName: ", u.firstName)
	fmt.Println("lastName: ", u.userLastName)
	fmt.Println("DOB: ", u.dob)
}

// This is a constructor function
func newUser(fname, lname, dob string) (*userInfo, error) {
	if fname == "" || lname == "" || dob == "" {
		return nil, errors.New("First name is empty or Last name is empty or Date of birth is empty")
	}
	return &userInfo{
		firstName:    fname,
		userLastName: lname,
		dob:          dob,
		createdAt:    time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserData("Enter user first name: ")
	userLastName := getUserData("Enter user last name: ")
	DOB := getUserData("Enter user data of birth: ")

	// user := userInfo{
	// 	firstName:    userFirstName,
	// 	userLastName: userLastName,
	// 	dob:          DOB,
	// 	createdAt:    time.Now(),
	// }

	user, err := newUser(userFirstName, userLastName, DOB)
	if err != nil {
		fmt.Println(err)
		return
	}

	PrintUserInfo(user)  // It's normal function to display the values of the structure
	user.PrintUserInfo() // It's method to a structure to display elements of the structure

}

// structure passing as value(pass by value)
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
	fmt.Scanln(&value)
	return value
}
