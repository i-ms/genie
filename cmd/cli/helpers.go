package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"os"
)

// setup Env loads the .env file if it exists
// setting RootPath for client application
// and database type
func setup() {
	err := godotenv.Load()
	if err != nil {
		exitGracefully(err)
	}

	path, err := os.Getwd()
	if err != nil {
		exitGracefully(err)
	}

	gen.RootPath = path
	gen.DB.DataType = os.Getenv("DATABASE_TYPE")
}

// getDSN build's the dsn for database connection and return dsn string
func getDSN() string {
	dbType := gen.DB.DataType

	// db driver being used is pgx so dbType is also pgx
	if dbType == "pgx" {
		dbType = "postgres"
	}

	if dbType == "postgres" {
		var dsn string
		if os.Getenv("DATABASE_PASS") != "" {
			dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_PASS"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"))
		} else {
			dsn = fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"))
		}
		return dsn
	} else {
		return "mysql://" + gen.BuildDSN()
	}
}

// showHelp displays the help menu
func showHelp() {
	color.Yellow(`Available commands:
help                    - show the help command
version                 - print application version
migrate                 - run database migrations
migrate down            - rollback most recent database migrations
migrate reset           - rollback all database migrations , and then run all migrations
make migration <name>   - creates two new up and down migrations in migration folder
make auth  			    - creates and runs migrations for authentication tables, and creates models and middleware    
make handler <name> 	- creates a stub handler in handler directory
make model <name> 		- creates a stub model in data directory
make session            - creates a table in database as a session store


`)
}
