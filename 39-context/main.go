package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Context is a built-in package in the Go standard library that provides a powerful toolset for managing concurrent operations.

context.Context is one of the MOST Important Concepts in Golang
This is critical in - production Go knowledge

Modern Go backend systems use: context.Context EVERYWHERE
It solves:
	cancellation
	timeouts
	deadlines
	request scoping
	tracing propagation
	graceful shutdown

Why Context Exists
Suppose:
	HTTP request arrives
	handler starts DB query
	DB query starts worker
	worker starts API call

Now client disconnects.

Question: how do ALL downstream operations stop?
Without cancellation:
	goroutines leak
	DB work wasted
	memory wasted
	resources wasted

Huge production problem.

Context Solves This. It provides: cancellation propagation tree

Context is: immutable + propagated explicitly. You PASS it everywhere.

Basic API
func Foo(ctx context.Context)
Note: context first argument by convention


Good Context Values
	request IDs
	auth info
	trace IDs
	correlation IDs

BAD Context Values
	database connection
	logger instance
	config object


*/

type contextKey string

const (
	RequestIDKey contextKey = "request_id"
	UserKey      contextKey = "user"
)

type User struct {
	ID       int
	Username string
}

func main() {
	var wg sync.WaitGroup

	// Simulate 5 Users requesting at the same time
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			// Root context
			ctx := context.Background()

			// Add request ID
			ctx = context.WithValue(
				ctx,
				RequestIDKey,
				fmt.Sprintf("req-%d", i),
			)

			// Simulated authenticated user
			ctx = context.WithValue(
				ctx,
				UserKey,
				User{
					ID:       i,
					Username: fmt.Sprintf("user-%d", i),
				},
			)

			handler(ctx)

		}(i)
	}

	wg.Wait()
}

func log(ctx context.Context, message string) {
	requestID := ctx.Value(RequestIDKey).(string)
	user := ctx.Value(UserKey).(User)

	fmt.Printf(
		"[RequestID=%s] [User=%s] %s\n",
		requestID,
		user.Username,
		message,
	)
}

func databaseLayer(ctx context.Context) error {
	log(ctx, "DB Query Started")

	time.Sleep(500 * time.Millisecond)

	requestID := ctx.Value(RequestIDKey).(string)

	// Simulate failure for one request
	if requestID == "req-3" {
		return fmt.Errorf("database connection failed")
	}

	log(ctx, "DB Query Finished")

	return nil
}

func serviceLayer(ctx context.Context) error {
	log(ctx, "Service Layer Started")

	err := databaseLayer(ctx)
	if err != nil {
		return fmt.Errorf(
			"service layer error: %w",
			err,
		)
	}

	log(ctx, "Service Layer Finished")

	return nil
}

func handler(ctx context.Context) {
	log(ctx, "Request Started")

	err := serviceLayer(ctx)
	if err != nil {
		log(ctx, fmt.Sprintf("ERROR: %v", err))

		return
	}

	log(ctx, "Request Completed Successfully")
}
