package auth

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
)

func GenerateSalt() (string, error) {
	saltBytes := make([]byte, 16) // 16 bytes = 128 bits
	_, err := rand.Read(saltBytes)
	if err != nil {
		return "", err
	}
	salt := base64.StdEncoding.EncodeToString(saltBytes)
	return salt, nil
}

func HashPassword(password string, salt string) string {
	saltBytes, _ := base64.StdEncoding.DecodeString(salt)
	passwordBytes := []byte(password)

	saltedPassword := append(passwordBytes, saltBytes...)
	hash := sha512.Sum512(saltedPassword)

	return base64.StdEncoding.EncodeToString(hash[:])
}

func verifyPassword(password string, salt string, hashedPassword string) bool {
	newHash := HashPassword(password, salt)
	return newHash == hashedPassword
}
