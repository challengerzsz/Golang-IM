package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	cipherStr := h.Sum(nil)
	encryptedStr := hex.EncodeToString(cipherStr)
	return strings.ToUpper(encryptedStr)
}

func ValidatePassword(plainPwd, salt, pwd string) bool {
	return Md5Encode(plainPwd+salt) == pwd
}

func MakePwd(plainPwd, salt string) string {
	return Md5Encode(plainPwd + salt)
}
