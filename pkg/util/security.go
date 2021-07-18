package util

import "golang.org/x/crypto/bcrypt"

// Hash return hash password
func Hash(password string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(password), 12) //goal: 16
	return string(h), err
}

// PasswordMatch return and error if match is't invalid
func PasswordMatch(hashed, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}
