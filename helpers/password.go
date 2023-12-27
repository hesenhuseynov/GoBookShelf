package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func CheckPassword(hashPassword string, plainPassword []byte) (bool, error) {
	hashPW := []byte(hashPassword)
	if err := bcrypt.CompareHashAndPassword(hashPW, plainPassword); err != nil {
		return false, err
	}

	return true, nil

}
