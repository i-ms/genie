package main

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"time"
)

func doAuth() error {

	// migrations
	dbType := gen.DB.DataType
	fileName := fmt.Sprintf("%d_create_auth_tables", time.Now().UnixMicro())
	upFile := gen.RootPath + "/migrations/" + fileName + ".up.sql"
	downFile := gen.RootPath + "/migrations/" + fileName + ".down.sql"

	log.Println(dbType, upFile, downFile)

	err := copyFileFromTemplate("templates/migrations/auth_tables."+dbType+".sql", upFile)
	if err != nil {
		exitGracefully(err)
	}

	err = copyDataToFile([]byte("DROP TABLE IF EXISTS users,tokens,remember_tokens CASCADE;"), downFile)
	if err != nil {
		exitGracefully(err)
	}

	// running migrations
	err = doMigrate("up", "")
	if err != nil {
		exitGracefully(err)
	}

	// copying files over
	err = copyFileFromTemplate("templates/data/user.go.txt", gen.RootPath+"/data/user.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/data/token.go.txt", gen.RootPath+"/data/token.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/middleware/auth.go.txt", gen.RootPath+"/middleware/auth.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/middleware/auth-token.go.txt", gen.RootPath+"/middleware/auth-token.go")
	if err != nil {
		exitGracefully(err)
	}

	color.Yellow(" - user, tokens and remember_tokens migration created successfully")
	color.Yellow(" - user and token models created")
	color.Yellow(" - auth middleware created")
	color.Yellow("")
	color.Yellow(" Don't forget to add user and token models in data/models, and to add appropriate middleware in routes!")

	return nil
}
