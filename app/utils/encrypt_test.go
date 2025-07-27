package utils

import (
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {
	t.Run("Encrypt and decrypt lowercase", func(t *testing.T) {
		text := "test"
		result, salt := Encrypt(text)
		decrypt, _ := Decrypt(result, salt)
		if decrypt != text {
			t.Log("Decrypted text must equal to original text")
		}
	})

	t.Run("Encrypt and decrypt uppercase", func(t *testing.T) {
		text := "TEST"
		result, salt := Encrypt(text)
		decrypt, _ := Decrypt(result, salt)
		if decrypt != text {
			t.Log("Decrypted text must equal to original text")
		}
	})

	t.Run("Encrypt and decrypt number", func(t *testing.T) {
		text := "12345"
		result, salt := Encrypt(text)
		decrypt, _ := Decrypt(result, salt)
		if decrypt != text {
			t.Log("Decrypted text must equal to original text")
		}
	})

	t.Run("Encrypt and decrypt mixed character", func(t *testing.T) {
		text := "sEcr3t"
		result, salt := Encrypt(text)
		decrypt, _ := Decrypt(result, salt)
		if decrypt != text {
			t.Log("Decrypted text must equal to original text")
		}
	})

	t.Run("Encrypt and decrypt special character", func(t *testing.T) {
		text := "sEcr3t!@129484weruoqrioeufp[q[woi[irwq[fdk]]]]"
		result, salt := Encrypt(text)
		decrypt, _ := Decrypt(result, salt)
		if decrypt != text {
			t.Log("Decrypted text must equal to original text")
		}
	})
}
