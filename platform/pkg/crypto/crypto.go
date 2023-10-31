package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"livecom/logger"
	"livecom/pkg/config"
)

func HashMD5(input string) string {
	// Create a new MD5 hasher
	hasher := md5.New()

	// Write the input string to the hasher
	hasher.Write([]byte(input))

	// Get the final hash as a byte array and convert it to a hexadecimal string
	hashedBytes := hasher.Sum(nil)
	hashedString := hex.EncodeToString(hashedBytes)

	return hashedString
}
func EncryptString( plaintext string) (string, error) {
	var secret = config.Cfg.EncryptionKey
	if(len(secret) < 32){
		logger.E(nil, "Error on EncryptString config.Cfg.EncryptionKey is too short - %s", secret)
		panic("Encryption key config.Cfg.EncryptionKey is too short")
	}
	return encrypt([]byte(secret), plaintext)
}
func DecryptString( plaintext string) (string, error) {
	var secret = config.Cfg.EncryptionKey
	if(len(secret) < 32){
		logger.E(nil, "Error on DecryptString  config.Cfg.EncryptionKey is too short - %s", secret)
		panic("Encryption key config.Cfg.EncryptionKey is too short")
	}
	return decrypt([]byte(secret), plaintext)
}
func encrypt(secretKey []byte, plaintext string) (string, error) {
    block, err := aes.NewCipher(secretKey)
    if err != nil {
        return "", err
    }

    // Create a random nonce (IV) for each encryption operation.
    nonce := make([]byte, 12)
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return "", err
    }

    aesGCM, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    ciphertext := aesGCM.Seal(nil, nonce, []byte(plaintext), nil)

    // Combine the nonce and ciphertext for later decryption.
    encryptedText := append(nonce, ciphertext...)

    // Encode the result in base64 to get a printable string.
    return base64.StdEncoding.EncodeToString(encryptedText), nil
}

func decrypt(secretKey []byte, encryptedText string) (string, error) {
    encryptedBytes, err := base64.StdEncoding.DecodeString(encryptedText)
    if err != nil {
        return "", err
    }

    block, err := aes.NewCipher(secretKey)
    if err != nil {
        return "", err
    }

    if len(encryptedBytes) < aes.BlockSize {
        return "", fmt.Errorf("ciphertext too short")
    }
	
    nonce := encryptedBytes[:12]
    ciphertext := encryptedBytes[12:]

    aesGCM, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return "", err
    }

    return string(plaintext), nil
}
