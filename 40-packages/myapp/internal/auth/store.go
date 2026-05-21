package auth

import (
	"errors"

	"myapp/pkg/utils"
)

var users = map[string]User{}

func CreateUser(
	email,
	password,
	username string,
) (*User, error) {

	if _, exists := users[email]; exists {
		return nil, errors.New("user already exists")
	}

	user := User{
		ID:       utils.GenerateRandomID(),
		Email:    email,
		Username: username,
		Password: password,
	}

	users[email] = user

	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	user, exists := users[email]

	if !exists {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
