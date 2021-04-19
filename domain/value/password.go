package value

import (
	"golang.org/x/crypto/bcrypt"
)

type Password string

func (pass Password) Digest() (PasswordDigest, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return PasswordDigest(hash), nil
}

type PasswordDigest string
