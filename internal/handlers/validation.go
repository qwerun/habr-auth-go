package handlers

import (
	"errors"
	"strings"
	"unicode"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (u *User) IsValid() error {
	if u.Email == "" || u.PasswordHash == "" || u.Nickname == "" {
		return errors.New("Missing required fields")
	}
	if err := u.validateEmail(); err != nil {
		return err
	}
	if err := u.validatePass(); err != nil {
		return err
	}
	if err := u.validateNick(); err != nil {
		return err
	}
	return nil
}

func (u *User) validateEmail() error {
	return validation.Validate(u.Email,
		validation.Required,
		is.Email,
	)
}

func (u *User) validateNick() error {
	nickname := u.Nickname
	if len(nickname) < 3 {
		return errors.New("Nickname is short (Minimum 3 characters)")
	}
	if len(nickname) > 25 {
		return errors.New("The nickname is too long")
	}
	allowedSumbols := "-_"
	var hasLetter, hasDigit bool
	for _, ch := range nickname {
		switch {
		case unicode.IsLower(ch), unicode.IsUpper(ch):
			hasLetter = true
		case unicode.IsDigit(ch):
			hasDigit = true
		case strings.ContainsRune(allowedSumbols, ch):

		default:
			return errors.New("The nickname can only contain Latin letters, numbers and symbols dashes and underlines")
		}
	}
	if !hasLetter || !hasDigit {
		return errors.New("The nickname must contain letters and numbers")
	}

	return nil
}

func (u *User) validatePass() error {
	password := u.PasswordHash
	if len(password) < 8 {
		return errors.New("Password is short (Minimum 8 characters)")
	}
	if len(password) > 64 {
		return errors.New("The password is too long")
	}
	allowedSumbols := "()-+=%\""
	var hasLetter, hasDigit bool
	for _, ch := range password {
		switch {
		case unicode.IsLower(ch), unicode.IsUpper(ch):
			hasLetter = true
		case unicode.IsDigit(ch):
			hasDigit = true
		case strings.ContainsRune(allowedSumbols, ch):

		default:
			return errors.New("The password can only contain Latin letters, numbers and symbols ()-+=%\"")
		}
	}
	if !hasLetter || !hasDigit {
		return errors.New("The password must contain letters and numbers")
	}

	return nil
}
