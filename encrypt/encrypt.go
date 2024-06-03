package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
)

func EncryptBytes(bytes []byte, key []byte) ([]byte, error) {
	key = ensureKeySize(key)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	copy(bytes, nonce)
	cipherText := gcm.Seal(nonce, nonce, bytes, nil)

	return cipherText, nil

}

func DecryptBytes(bytes []byte, key []byte) ([]byte, error) {
	key = ensureKeySize(key)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := bytes[:gcm.NonceSize()]
	cipherText := bytes[gcm.NonceSize():]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil

}
func ensureKeySize(key []byte) []byte {
	hash := sha256.New()
	hash.Write(key)
	return hash.Sum(nil)[:32] // Returning the first 32 bytes for AES-256
}
