package cryptx

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	salt = "(momo)@#$"
)

// 哈希，纯文本
func Hash(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedBytes), err
}

// 比较，密文和明文是否相同
func Compare(hashed, plainText string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plainText))
}

func MD5(plainText string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(plainText + salt)))
}

// Base64 编码
func Encode(src string) string {
    return base64.StdEncoding.EncodeToString([]byte(src))
}

// Base64 解码
func Decode(src string) (string, error) {
    originalData, err := base64.StdEncoding.DecodeString(src)
    if err != nil {
        return "", err
    }
    return string(originalData), nil
}

// Tips - 命名差异，util HashPassword\CheckPassword 需要强调用途
