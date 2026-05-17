package main

import "fmt"

/*
Enums in Go
Go does NOT have traditional enums like: java, c++, typescript
Go instead uses constants + custom types

What Is iota?
iota is Go’s special auto-increment identifier.
Inside const block:
(
	1,
	2,
	3,
	...
)
*/

type Status int

const (
	Pending Status = iota
	Processing
	Completed
	Failed
)

func (s Status) String() string {
	switch s {
	case Pending:
		return "pending"
	case Processing:
		return "processing"
	case Completed:
		return "completed"
	case Failed:
		return "failed"
	default:
		return "unknown"
	}
}

// String-Based Enums
type Role string

const (
	Admin Role = "ADMIN"
	User  Role = "USER"
	Guest Role = "GUEST"
)

func main() {
	var status Status = Processing
	fmt.Println(status) // without the String() method, output -> 1 | with the String() method, output -> processing
	/*
		Because:
		fmt checks for String() method
		uses fmt.Stringer interface
	*/

	fmt.Println(Admin) // Admin
}
