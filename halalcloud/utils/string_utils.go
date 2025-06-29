package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func IsEmptyString(s string) bool {
	return s == "" || s == "null" || s == "undefined"
}

func CreateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 使用crypto/rand包生成更安全的随机字符串
	if length <= 0 {
		return ""
	}
	if length > 256 {
		length = 256 // 限制最大长度为256
	}
	result := make([]byte, length)
	_, err := rand.Read(result)
	if err != nil {
		return ""
	}
	for i := 0; i < length; i++ {
		result[i] = charset[result[i]%byte(len(charset))]
	}

	// 返回生成的随机字符串
	// 这里使用了crypto/rand包来生成更安全的随机字节
	// 然后将其转换为指定长度的字符串
	return string(result)
}

func Sha256Hash(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

func Sha256HashString(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
