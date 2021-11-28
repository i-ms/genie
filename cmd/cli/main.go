package main

import (
	"errors"
	"github.com/fatih/color"
	"github.com/i-ms/genie"
	"log"
	"os"
)

const version = "1.0.0"

var gen genie.Genie

func main() {
	arg1, arg2, arg3, err := validateInput()
	if err != nil {
		exitGracefully(err)
	}

	switch arg1 {
	case "help":
		showHelp()

	case "version":
		showVersion()

	default:
		log.Println(arg2, arg3)
	}
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
			arg3 = os.Args[4]
		}
	} else {
		color.Red("Error: No arguments provided")
		showHelp()
		return "", "", "", errors.New("no command provided")
	}

	return arg1, arg2, arg3, nil
}

func showHelp() {
	color.Yellow(`Available commands
    help        -   show the help command
    version     -   print application version`)
}

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

func showVersion() {
	color.Yellow("Application version: %v", version)
	return
}
