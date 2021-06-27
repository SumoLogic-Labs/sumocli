package encryption

import "testing"

func TestEncryption(t *testing.T) {
	secret := "secretstring"
	encryptedText := EncryptData(secret)
	decryptedText := DecryptData(encryptedText)
	if decryptedText != secret {
		t.Errorf("Encryption and decryption failed, got: %s, expected: %s", decryptedText, secret)
	}
}
