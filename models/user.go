package models

import(
	"errors"
)

// User is a model of the user
type User struct{
	ID int
	FirstName string
	LastName string
}

var (
	users []*User
	nextID = 1
)

// GetUsers gets all users of the model
func GetUsers() []*User {
	return users
}

// AddUser add user to model
func AddUser(u User)  (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New user must not include id or it must be set to zero")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

// GetUserByID returns user given ID
func GetUserByID(id int) (User, error){
	
	return User{}, nil
}