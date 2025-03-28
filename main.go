package main

import (
	"fmt"
	"os"

	"github.com/wfcornelissen/tyrecheck/internal/cmd"
	"github.com/wfcornelissen/tyrecheck/internal/dbFuncs"
)

func main() {
	cmd.Execute()
	// Read .env file. If firstrun is true, create tables
	firstrun := os.Getenv("FIRST_RUN")
	if firstrun == "true" {
		fmt.Println("First run detected from .env file")
		// Check whether tyrecheck.db exists
		if _, err := os.Stat("./tyrecheck.db"); os.IsNotExist(err) {
			fmt.Println("tyrecheck.db does not exist, creating tables")
			dbFuncs.CreateTables()
		} else {
			fmt.Println("tyrecheck.db exists, skipping table creation")
		}
		// Set firstrun to false
		os.Setenv("FIRST_RUN", "false")
	}

}
