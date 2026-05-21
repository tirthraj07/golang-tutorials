package main

import (
	"fmt"

	"myapp/internal/auth"
	"myapp/internal/logger"
	"myapp/internal/user"
)

func main() {
	logger.Info("Application Started")

	// Signup
	newUser, err := auth.Signup(
		"tirthraj@gmail.com",
		"password123",
		"tirthraj07",
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("User Created:", newUser)

	// Login
	session, err := auth.Login(
		"tirthraj@gmail.com",
		"password123",
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Session:", session)

	// Validate Session
	currentUser, err := auth.GetSessionUser(session.Token)

	if err != nil {
		panic(err)
	}

	fmt.Println("Current User:", currentUser)

	// User package
	profile := user.GetProfile(currentUser.ID)

	fmt.Println("Profile:", profile)

	preferences := user.GetPreferences(currentUser.ID)

	fmt.Println("Preferences:", preferences)

	logger.Info("Application Finished")
}
