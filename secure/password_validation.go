package secure

import "regexp"

func IsStrongPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	lowercase, _ := regexp.MatchString(`[a-z]`, password)
	uppercase, _ := regexp.MatchString(`[A-Z]`, password)
	digit, _ := regexp.MatchString(`[0-9]`, password)
	specialChar, _ := regexp.MatchString(`[!@#\$%\^&*]`, password)

	return lowercase && uppercase && digit && specialChar
}
