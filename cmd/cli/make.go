package main

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"io/ioutil"
	"strings"
	"time"
)

func doMake(arg2, arg3 string) error {
	switch arg2 {
	case "migration":
		makeMigration(arg3)

	case "auth":
		err := doAuth()
		if err != nil {
			exitGracefully(err)
		}

	case "handler":
		handler(arg3)

	case "helper":
		helper(arg3)

	case "model":
		model(arg3)

	case "session":
		err := doSessionTable()
		if err != nil {
			exitGracefully(err)
		}
	}

	return nil
}

func makeMigration(arg3 string) {
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
}

func handler(arg3 string) {
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

func model(arg3 string) {
	if arg3 == "" {
		exitGracefully(errors.New("please specify model name"))
	}

	// read the model template in data
	data, err := templateFS.ReadFile("templates/data/model.go.txt")
	if err != nil {
		exitGracefully(err)
	}

	// converting data to string
	// for utilizing it with pluralize
	model := string(data)

	plur := pluralize.NewClient()

	var modelName = arg3
	var tableName = arg3

	if plur.IsPlural(arg3) {
		modelName = plur.Singular(arg3)
		tableName = strings.ToLower(tableName)
	} else {
		tableName = strings.ToLower(plur.Plural(arg3))
	}

	fileName := gen.RootPath + "/data/" + strings.ToLower(modelName) + ".go"
	if fileExists(fileName) {
		exitGracefully(errors.New(fileName + " model already exists"))
	}

	// replacing placeholder with user specified values
	model = strings.ReplaceAll(model, "$MODELNAME$", strcase.ToCamel(modelName))
	model = strings.ReplaceAll(model, "$TABLENAME$", tableName)

	// writing data to file
	err = copyDataToFile([]byte(model), fileName)
	if err != nil {
		exitGracefully(err)
	}
}

func helper(arg3 string) {
	if arg3 == "" {
		exitGracefully(errors.New("please specify type of helper function to be created"))
	}

	switch arg3 {
	case "handler":
		handlerHelper()
	case "routes":
		routesHelper()
	default:
		exitGracefully(errors.New("invalid helper type"))
	}
}

func handlerHelper() {
	var fileName string
	color.Blue("Name of file to be created: ")
	fmt.Scanln(&fileName)
	targetAddr := fmt.Sprintf("%s/handlers/%s.go", gen.RootPath, fileName)

	if fileExists(targetAddr) {
		exitGracefully(errors.New(fileName + " already exists"))
	}

	err := copyFileFromTemplate("templates/handlers/handler-helper.go.txt", targetAddr)
	if err != nil {
		exitGracefully(err)
	}
}

func routesHelper() {
	var fileName string
	color.Blue("Name of file to be created: ")
	fmt.Scanln(&fileName)
	targetAddr := fmt.Sprintf("%s/%s.go", gen.RootPath, fileName)
	
	if fileExists(targetAddr) {
		exitGracefully(errors.New(fileName + " already exists"))
	}
	
	err:= copyFileFromTemplate("templates/routes/routes-helper.go.txt", targetAddr)
	if err != nil {
		exitGracefully(err)
	}
}
