package genie

import (
	"crypto/rand"
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
