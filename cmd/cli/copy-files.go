package main

import (
	"embed"
	"io/ioutil"
)

//go:embed templates
var templateFS embed.FS

func copyFileFromTemplate(templatePath string, targetFile string) error {
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
