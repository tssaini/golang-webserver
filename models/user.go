package models

import(
	"errors"
	"fmt"
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
	for _, u := range users{
		if u.ID == id{
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}
// UpdateUser update an existing user
func UpdateUser(u User) (User, error){
	for i, user := range users{
		if user.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

// RemoveUserByID remove a user with a given ID
func RemoveUserByID(id int) error{
	for i, user := range users{
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)
}