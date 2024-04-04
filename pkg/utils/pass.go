package utils

import (
	"errors"
	"regexp"
	"strings"
)

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password length must be at least 8 characters")
	}

	// Check for uppercase letter
	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return errors.New("password must contain at least one uppercase letter")
	}

	// Check for lowercase letter
	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		return errors.New("password must contain at least one lowercase letter")
	}

	// Check for numeric digit
	if !strings.ContainsAny(password, "0123456789") {
		return errors.New("password must contain at least one numeric digit")
	}

	// Check for special character
	re := regexp.MustCompile(`[!@#$%^&*()-_+=\[\]{}|:;"'<>,.?/~]`)
	if !re.MatchString(password) {
		return errors.New("password must contain at least one special character")
	}

	return nil
}
