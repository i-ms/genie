package main

import (
	"github.com/fatih/color"
)

func doMigrate(arg2, arg3 string) error {
	dsn := getDSN() 

	// running migrations command
	switch arg2 {
	case "up":
		err := gen.MigrateUp(dsn)
		if err != nil {
			return err
		}
	case "down":
		if arg3 == "all" {
			err := gen.MigrateDownAll(dsn)
			if err != nil {
				return err
			}
		} else {
			err := gen.Steps(-1, dsn)
			if err != nil {
				return err
			}
		}
	case "reset":
		// perform migration down : removing all migrations from db
		err := gen.MigrateDownAll(dsn)
		if err != nil {
			return err
		}
		// perform migration up : running all migrations
		err = gen.MigrateUp(dsn)
		if err != nil {
			return err
		}
	default:
		color.Red("Unknown command")
		showHelp()
	}
	return nil
}
