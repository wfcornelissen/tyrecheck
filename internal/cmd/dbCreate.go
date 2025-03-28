/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wfcornelissen/tyrecheck/internal/dbFuncs"
)

// dbCreateCmd represents the dbCreate command
var dbCreateCmd = &cobra.Command{
	Use:   "dbCreate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat("./tyrecheck.db"); os.IsNotExist(err) {
			fmt.Println("tyrecheck.db does not exist, creating tables")
			dbFuncs.CreateTables()
		} else {
			fmt.Println("tyrecheck.db exists, skipping table creation")
		}
	},
}

func init() {
	rootCmd.AddCommand(dbCreateCmd)

}
