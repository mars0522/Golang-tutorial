package main

import (
	"concurrency/Basics/user"
	"fmt"
)

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

	employee, err := user.NewUser(userFirstName, userLastName, DOB)
	if err != nil {
		fmt.Println(err)
		return
	}
	PrintUserInfo(employee)  // It's normal function to display the values of the structure
	employee.PrintUserInfo() // It's method to a structure to display elements of the structure

	admin, err := user.CreateAdmin("varun@gmail.com", "password123", "Varun", "Singh", "1990-01-01")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Printting admin info
	admin.User.PrintUserInfo()

}

// structure passing as value(pass by value)
func PrintUserInfo(user *user.UserInfo) {
	fmt.Printf("User Info:\n")
	fmt.Printf("First Name: %s\n", user.FirstName)
	fmt.Printf("Last Name: %s\n", user.UserLastName)
	fmt.Printf("Date of Birth: %s\n", user.Dob)
	fmt.Printf("Created At: %s\n", user.CreatedAt)
	fmt.Printf("-----------------------------\n")
}
func getUserData(promtText string) string {
	fmt.Print(promtText)
	var value string
	fmt.Scanln(&value)
	return value
}
