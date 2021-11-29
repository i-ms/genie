package main

import (
	"embed"
	"errors"
	"io/ioutil"
	"os"
)

//go:embed templates
var templateFS embed.FS

// copy contents of one file to another
func copyFileFromTemplate(templatePath string, targetFile string) error {
	if fileExists(targetFile) {
		return errors.New(targetFile + " already exists")
	}

	data, err := templateFS.ReadFile(templatePath)
	if err != nil {
		exitGracefully(err)
	}

	err = copyDataToFile(data, targetFile)
	if err != nil {
		exitGracefully(err)
	}

	return nil
}

// copyDataToFile writes specified data to target file
func copyDataToFile(data []byte, to string) error {
	err := ioutil.WriteFile(to, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
