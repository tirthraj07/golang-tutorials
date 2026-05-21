package auth

import (
	"errors"

	"myapp/pkg/utils"
)

var sessions = map[string]Session{}

func CreateSession(userID string) (*Session, error) {
	token := utils.GenerateRandomID()

	session := Session{
		Token:  token,
		UserID: userID,
	}

	sessions[token] = session

	return &session, nil
}

func ValidateSession(token string) (*Session, error) {
	session, exists := sessions[token]

	if !exists {
		return nil, errors.New("invalid session")
	}

	return &session, nil
}
