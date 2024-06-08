package lib

import (
  "crypto/md5"
	"crypto/aes"
	"crypto/rand"
	"crypto/cipher"
	"encoding/hex"
	"io/ioutil"
	"errors"
	"io"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) ([]byte, error) {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return []byte(""), errors.New("Error when encrypting the data.")
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

func decrypt(data []byte, passphrase string) ([]byte, error) {
	key := []byte(createHash(passphrase))
	block, cipherErr := aes.NewCipher(key)
	gcm, gcmErr := cipher.NewGCM(block)
	if cipherErr != nil || gcmErr != nil {
		return []byte(""), errors.New("Error when decrypting a file.")
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return []byte(""), errors.New("Error when decrypting a file.")
	}
	return plaintext, nil
}

func EncryptWrite(lockername string, value string) error {
	key, ok := keys[lockername]
	filename := LockerPath + "/" + lockername + ".aes"
	if ok {
  	cipher, err := encrypt([]byte(value), key)
		ioutil.WriteFile(filename, cipher, 777)
  	return err
	}
	return errors.New("No key submitted for this locker.")
}

func DecryptRead(filename string, key string) (string, error) {
	data, readErr := ioutil.ReadFile(filename)
	if readErr != nil {
		return "", errors.New("Error when reading encryption file.")
	}
  plaintext, decryptErr := decrypt(data, key)
  if decryptErr != nil || len(plaintext) == 0 {
		return "", errors.New("Error when decrypting file contents.")
  }
	return string(plaintext), nil
}

