package user

func GetProfile(userID string) map[string]string {
	return map[string]string{
		"user_id": userID,
		"bio":     "Backend Engineer",
		"country": "India",
	}
}
