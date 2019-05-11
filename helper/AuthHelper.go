package helper

import "golang.org/x/crypto/bcrypt"

func IsSame(str string, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(str)) == nil
}

func Hash(str *string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(*str), bcrypt.DefaultCost)
	return string(hashed), err
}
