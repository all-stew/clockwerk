package password_util

import (
	"clockwerk/pkg/util/string_util"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func GetSalt() string {
	return string_util.RandStr(32)
}

func GetPassword(password string, salt string) string {
	passwordWithSlat := fmt.Sprintf("%s/%s", password, salt)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(passwordWithSlat), bcrypt.DefaultCost)
	return string(hashedPassword)
}
