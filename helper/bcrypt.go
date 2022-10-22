package helper

import "golang.org/x/crypto/bcrypt"

func HashPass(p string) string {
	salt := 9
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)
	return string(hash)
}
