package user

func GetPreferences(userID string) map[string]string {
	return map[string]string{
		"theme":    "dark",
		"language": "en",
	}
}
