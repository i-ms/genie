package main

import (
	"errors"
	"fmt"
	"github.com/iancoleman/strcase"
	"io/ioutil"
	"strings"
	"time"
)

func doMake(arg2, arg3 string) error {
	switch arg2 {
	case "migration":
		dbType := gen.DB.DataType
		if arg3 == "" {
			exitGracefully(errors.New("please specify migration name"))
		}
		fileName := fmt.Sprintf("%d_%s", time.Now().UnixMicro(), arg3)

		upFile := gen.RootPath + "/migrations/" + fileName + "." + dbType + ".up.sql"
		downFile := gen.RootPath + "/migrations/" + fileName + "." + dbType + ".down.sql"

		err := copyFileFromTemplate("templates/migrations/migration."+dbType+".up.sql", upFile)
		if err != nil {
			exitGracefully(err)
		}

		err = copyFileFromTemplate("templates/migrations/migration."+dbType+".down.sql", downFile)
		if err != nil {
			exitGracefully(err)
		}

	case "auth":
		err := doAuth()
		if err != nil {
			exitGracefully(err)
		}

	case "handler":
		if arg3 == "" {
			exitGracefully(errors.New("please specify handler name"))
		}

		fileName := gen.RootPath + "/handlers/" + strings.ToLower(arg3) + ".go"

		if fileExists(fileName) {
			exitGracefully(errors.New(fileName + " handler already exists"))
		}

		data, err := templateFS.ReadFile("templates/handlers/handler.go.txt")
		if err != nil {
			exitGracefully(err)
		}
		
		handler := string(data)
		handler = strings.ReplaceAll(handler, "$HANDLERNAME$", strcase.ToCamel(arg3))

		err = ioutil.WriteFile(fileName, []byte(handler), 0644)
		if err != nil {
			exitGracefully(err)
		}

	}
	
	return nil
}
