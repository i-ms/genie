package genie

import "os"

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
