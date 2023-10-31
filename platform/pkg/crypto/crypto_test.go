package crypto

import (
	"testing"
)

func TestEncryptionAndDecryption(t *testing.T) {
	// Sample secretKey and plaintext.
	secretKey := []byte("81d1070ee37c70cd73fdf9b592f67a4d") // 32 bytes for AES-256
	plaintext := "A random hexadecimal number generator can be useful if you're doing cross-browser testing. For example, you can generate random MD5 hashes (hex numbers of length 32) or random SHA1 git hashes (hex numbers of length 40). These values can be used as unique identifiers for cached files or temporary resources as the likelihood of long random hex value collisions is very low. Similarly, you can generate random hex numbers of a certain length to enter in forms to test form validation code, as well as use random hexadecimal values as random data.	Looking for more web developer tools? Try these!"

	// Encrypt the plaintext.
	encryptedText, err := encrypt(secretKey, plaintext)

	if err != nil {
		t.Fatalf("Failed to encrypt: %s", err)
	}

	if encryptedText == plaintext {
		t.Fatalf("Encrypted text is the same as plaintext!")
	}

	// Decrypt the encrypted text.
	
	decryptedText, err := decrypt(secretKey, encryptedText)

	if err != nil {
		t.Fatalf("Failed to decrypt: %s", err)
	}
	t.Logf("encryptedText: %v",encryptedText)
	// t.Logf("Time taken to decrypt: %v", elapsedDecrypt)
	if decryptedText != plaintext {
		t.Fatalf("Decrypted text (%s) does not match original plaintext (%s)", decryptedText, plaintext)
	}
}

// Additional tests can be added as needed.
