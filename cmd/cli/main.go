package main

import (
	"errors"
	"github.com/fatih/color"
	"github.com/i-ms/genie"
	"os"
)

const version = "1.0.0"

var gen genie.Genie

func main() {
	var message string
	arg1, arg2, arg3, err := validateInput()
	if err != nil {
		exitGracefully(err)
	}

	setup()

	switch arg1 {
	case "help":
		showHelp()

	case "version":
		showVersion()

	case "migrate":
		if arg2 == "" {
			arg2 = "up"
		}
		err = doMigrate(arg2, arg3)
		if err != nil {
			exitGracefully(err)
		}
		message = "Migrations executed successfully"

	case "make":
		if arg2 == "" {
			exitGracefully(errors.New("make requires a subcommand : (migration | model | handler)"))
		}
		if err = doMake(arg2, arg3); err != nil {
			exitGracefully(err)
		}

	default:
		color.Red("Unknown command")
		showHelp()
	}

	exitGracefully(nil, message)
}

// validateInput validates the input arguments
func validateInput() (string, string, string, error) {
	var arg1, arg2, arg3 string

	if len(os.Args) > 1 {
		arg1 = os.Args[1]

		if len(os.Args) >= 3 {
			arg2 = os.Args[2]
		}

		if len(os.Args) >= 4 {
			arg3 = os.Args[3]
		}
	} else {
		color.Red("Error: No arguments provided")
		showHelp()
		return "", "", "", errors.New("No command provided")
	}

	return arg1, arg2, arg3, nil
}

// exitGracefully exits the application and provides the error or optional message
func exitGracefully(err error, msg ...string) {
	message := ""
	if len(msg) > 0 {
		message = msg[0]
	}

	if err != nil {
		color.Red("Error: %v\n", err)
	}

	if len(message) > 0 {
		color.Yellow(message)
	} else {
		color.Green("Finished!")
	}

	os.Exit(0)
}

// showVersion displays the current application version
func showVersion() {
	color.Yellow("Application version: %v", version)
	return
}
