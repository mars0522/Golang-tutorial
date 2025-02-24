package user

import (
	"errors"
	"fmt"
	"time"
)

type UserInfo struct {
	FirstName    string
	UserLastName string
	Dob          string
	CreatedAt    time.Time
}

type Admin struct {
	Email    string
	Password string
	User     UserInfo
}

// This is a go Method
func (u *UserInfo) PrintUserInfo() {
	fmt.Println("firstName: ", u.FirstName)
	fmt.Println("lastName: ", u.UserLastName)
	fmt.Println("DOB: ", u.Dob)
}

func CreateAdmin(email, password, firstName, lastName, dob string) (*Admin, error) {
	if email == "" || password == "" || firstName == "" || lastName == "" || dob == "" {
		return nil, errors.New("email, password, first name, last name, and date of birth cannot be empty")
	}

	return &Admin{
		Email:    email,
		Password: password,
		User:     UserInfo{FirstName: firstName, UserLastName: lastName, Dob: dob, CreatedAt: time.Now()},
	}, nil
}

// This is a constructor function
func NewUser(fname, lname, dob string) (*UserInfo, error) {
	if fname == "" || lname == "" || dob == "" {
		return nil, errors.New("first name is empty or Last name is empty or Date of birth is empty")
	}
	return &UserInfo{
		FirstName:    fname,
		UserLastName: lname,
		Dob:          dob,
		CreatedAt:    time.Now(),
	}, nil
}
