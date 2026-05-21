package utils

import (
	"fmt"
	"math/rand"
)

func GenerateRandomID() string {
	return fmt.Sprintf(
		"id-%d",
		rand.Intn(100000),
	)
}
