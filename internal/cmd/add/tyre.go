/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package add

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wfcornelissen/tyrecheck/internal/entries"
)

// tyreCmd represents the tyre command
var tyreCmd = &cobra.Command{
	Use:   "tyre",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tyre called")
		entries.AddTyre()
	},
}

func init() {
	AddCmd.AddCommand(tyreCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tyreCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tyreCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
