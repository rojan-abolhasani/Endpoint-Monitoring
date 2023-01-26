package util

import (
	"monitor/model"
	"strings"
	"unicode"
)

// validates the password and the username, returns an error in case of a problem
func ValidateUser(s model.RegisterUserRequest) bool {
	if !ValidatePassword(*s.PassWord) || !ValidateUserName(*s.UserName) {
		return false
	}
	return true
}

// Username conditions: not less than 5 characters, the first character should be a letter, only numbers and letters are allowed
func ValidateUserName(u string) bool {
	if len(u) < 5 {
		return false
	}
	if !unicode.IsLetter(rune(u[0])) {
		return false
	}
	for _, v := range u {
		if !unicode.IsLetter(v) && !unicode.IsDigit(v) {
			return false
		}
	}
	return true
}

// conditions for the password : length must be more than 10 characters, at least one digit, lowercase letter, uppercase letter and a character
func ValidatePassword(u string) bool {
	has_digit := false
	has_lower := false
	has_upper := false
	has_charachter := false
	if len(u) < 10 {
		return false
	}
	for _, v := range u {
		if unicode.IsDigit(v) {
			has_digit = true
		} else if unicode.IsLower(v) {
			has_lower = true
		} else if unicode.IsUpper(v) {
			has_upper = true
		} else {
			has_charachter = true
		}
	}
	if !has_digit || !has_charachter || !has_lower || !has_upper {
		return false
	}
	return true
}

// methods can only be options, head and get
func ValidateMethod(u string) bool {
	m := strings.ToUpper(u)
	if m == "OPTIONS" || m == "HEAD" || m == "GET" {
		return true
	}
	return false
}
