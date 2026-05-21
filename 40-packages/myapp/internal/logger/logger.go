package logger

import (
	"fmt"
	"time"
)

func Info(message string) {
	fmt.Printf(
		"[INFO] [%v] %s\n",
		time.Now().Format(time.RFC3339),
		message,
	)
}
