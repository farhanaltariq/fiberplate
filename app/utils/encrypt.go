package utils

import (
	"encoding/base64"
	"errors"
	"math/rand"
)

func GenerateRandomNumber(min, max int) int {
	for {
		value := rand.Intn(max-min+1) + min
		if value >= min && value <= max {
			return value
		}
	}
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = charset[rand.Intn(len(charset))]
	}
	return string(randomString)
}

func GenerateSalt() string {
	len := GenerateRandomNumber(4, 8)
	return GenerateRandomString(len)
}

func Encrypt(password string) (string, string) {
	salt := GenerateSalt()
	signature := GetEnv("SIGNATURE_KEY", "secret")

	sigIndex := 0
	encPassword := make([]rune, 0)
	for _, v := range password {
		if sigIndex == len(signature) {
			sigIndex = 0
		}

		encPassword = append(encPassword, v)
		encPassword = append(encPassword, rune(signature[sigIndex]))
		sigIndex++
	}

	sigIndex = 0
	for i := range encPassword {
		if sigIndex == len(salt) {
			sigIndex = 0
		}

		encPassword[i] += rune(salt[sigIndex])

		sigIndex++
	}

	stringEncPassword := base64.StdEncoding.EncodeToString([]byte(string((encPassword))))
	return stringEncPassword, salt
}

func Decrypt(password string, salt string) (string, error) {
	// signature := GetEnv("SIGNATURE_KEY", "secret")

	plainEnc, err := base64.StdEncoding.DecodeString(password)
	if err != nil {
		return "", errors.New("failed to decode password")
	}

	sigIndex := 0
	pass := []rune(string(plainEnc))
	for i := range pass {
		if sigIndex == len(salt) {
			sigIndex = 0
		}

		pass[i] -= rune(salt[sigIndex])

		sigIndex++
	}

	sigIndex = 0
	decodedPass := make([]rune, 0)
	for i := range pass {
		if i%2 == 0 {
			decodedPass = append(decodedPass, pass[i])
		}
	}

	return string(decodedPass), nil
}
