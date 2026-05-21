package auth

type User struct {
	ID       string
	Email    string
	Username string
	Password string
}

type Session struct {
	Token  string
	UserID string
}
