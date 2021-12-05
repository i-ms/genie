package genie

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"
)

const (
	randomString = "qwertyuioplkjhgfdaszxcvbnmQWERTYUIOPLKJHBGVFCDXSAZNM1234567890+-"
)

// RandomString generates a random string length n from values in the const randomString
func (g *Genie) RandomString(n int) string {
	s, r := make([]rune, n), []rune(randomString)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}
	return string(s)
}

// CreateDirIfNotExists : creates directory
// If there will be error , it will be off type error
func (g *Genie) CreateDirIfNotExists(path string) error {
	const mode = 0755
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, mode)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateFileIfNotExists : creates a new file with specified name
func (g *Genie) CreateFileIfNotExists(file string) error {
	var _, err = os.Stat(file)
	if os.IsNotExist(err) {
		var file, err = os.Create(file)
		if err != nil {
			return err
		}

		// Closing file
		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	}
	return nil
}

type Encryption struct {
	Key []byte
}

// Encrypt takes plain text as input and provides the encrypted text
// along with error if present
func (e *Encryption) Encrypt(text string) (string, error) {
	plainText := []byte(text)

	block, err := aes.NewCipher(e.Key)
	if err != nil {
		return "", err
	}

	// data structure to be used for encryption
	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plainText)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt convert's encrypted text to plain text and provides error ( if present )
func (e *Encryption) Decrypt(cryptoText string) (string, error) {
	cipherText, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(e.Key)
	if err != nil {
		return "", nil
	}

	if len(cipherText) < aes.BlockSize {
		return "", err
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}
