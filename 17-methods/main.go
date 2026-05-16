package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"regexp"
)

/*
Methods are just functions attached to types

Important Syntax

func (t Type) foo(){}

This part - (t Type) - is called a 'receiver'. It means that the method belongs to the User type

Methods are not INSIDE structs, they are declared separately

Note:
(t Type) ->  receives COPY of struct. Exactly like normal function arguments.
(t *Type) -> receives the struct itself.

Go has visibility rules:
	uppercase = exported/public
	lowercase = package-private

Go does NOT enforce OOP-style encapsulation heavily.

It prefers:
	package boundaries
	simple APIs
over:
	deep inheritance trees
	private/protected complexity

Go does NOT have constructors. This is another intentional Go design decision.
Go avoids:
	special object lifecycle rules
	hidden initialization behavior
	constructor overloading complexity

Instead, Go uses: normal functions -> Usually named: NewType

Getter is usually named: same as field NOT GetName()
Setters is usually named: SetName()
Go prefers:
	user.Name()
NOT:
	user.GetName()

Go developers often simply do: user.Name = "Tirthraj"
No getter/setter. This is considered perfectly fine in Go.

In Java, we have getters and setters even when methods do nothing special

Encapsulation Used Only When Needed

Examples:
	validation
	computed values
	synchronization
	invariants
	lazy loading

*/

type User struct {
	userId   string
	username string
	email    string
}

func NewUser(username, email string) (*User, error) {
	userId, err := generateRandomId(10)
	if err != nil {
		return nil, fmt.Errorf("unable to create User: %w", err)
	}
	user := User{}
	if err := user.setUserId(userId); err != nil {
		return nil, err
	}

	if err := user.setUsername(username); err != nil {
		return nil, err
	}

	if err := user.setEmail(email); err != nil {
		return nil, err
	}

	return &user, nil
}

// Regex
var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`)
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
var userIdRegex = regexp.MustCompile(`^[a-zA-Z0-9]{10}$`)

// Methods
func (u *User) setUsername(username string) error {
	if !usernameRegex.MatchString(username) {
		return errors.New("Invalid Username")
	}
	u.username = username
	return nil
}

func (u *User) setEmail(email string) error {
	if !emailRegex.MatchString(email) {
		return errors.New("Invalid Email")
	}

	u.email = email
	return nil
}

func (u *User) setUserId(userId string) error {
	if !userIdRegex.MatchString(userId) {
		return errors.New("Invalid UserId")
	}

	u.userId = userId
	return nil
}

// Helper Functions (Normal Function)
func generateRandomId(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	const maxByte = 256
	maxValidByte := maxByte - (maxByte % len(charset))

	result := make([]byte, length)

	randomBytes := make([]byte, length)

	i := 0

	for i < length {
		_, err := rand.Read(randomBytes)

		if err != nil {
			return "", fmt.Errorf("unable to generate random string: %w", err)
		}

		for _, b := range randomBytes {
			if i >= length {
				break
			}

			if int(b) >= maxValidByte {
				continue
			}

			result[i] = charset[int(b)%len(charset)]

			i++
		}
	}

	return string(result), nil
}

func main() {
	user, err := NewUser("tirthraj07", "tirthraj@gmail.com")
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %w", err))
		return
	}

	fmt.Println(*user)
}
