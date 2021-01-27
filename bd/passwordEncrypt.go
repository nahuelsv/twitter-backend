package bd

import "golang.org/x/crypto/bcrypt"

/*PasswordEncrypt function that returns encrypted password*/
func PasswordEncrypt(pass string) (string, error) {
	steps := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), steps)
	return string(bytes), err
}
