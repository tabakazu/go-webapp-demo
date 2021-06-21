package value

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type UsernameOrEmail string

func (v UsernameOrEmail) IsEmail() bool {
	re := regexp.MustCompile(`[a-zA-Z0-9]+@.+\..`)
	return re.FindString(string(v)) != ""
}

type Password string

func (v Password) Digest() (PasswordDigest, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(v), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return PasswordDigest(b), nil
}

type PasswordDigest string

func (v PasswordDigest) ValidPassword(passwd Password) error {
	if err := bcrypt.CompareHashAndPassword([]byte(v), []byte(passwd)); err != nil {
		return err
	}
	return nil
}
