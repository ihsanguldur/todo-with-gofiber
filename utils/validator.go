package utils

import "net/mail"

func IsPasswordValid(password string) bool {
	return len(password) >= 6
}

func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil

}
