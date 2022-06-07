package utils

import (
	"golang.org/x/crypto/bcrypt"
)

/*
   说明：用户密码加密，验证比对
*/

// CryptoPassword 生成加密密码
func CryptoPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

// ComparePassword 密码对比校验
func ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
