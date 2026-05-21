package auth

import "errors"

func Signup(
	email,
	password,
	username string,
) (*User, error) {

	user, err := CreateUser(
		email,
		password,
		username,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func Login(
	email,
	password string,
) (*Session, error) {

	user, err := GetUserByEmail(email)

	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, errors.New("invalid password")
	}

	return CreateSession(user.ID)
}

func GetSessionUser(token string) (*User, error) {
	session, err := ValidateSession(token)

	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.ID == session.UserID {
			return &user, nil
		}
	}

	return nil, errors.New("user not found")
}
