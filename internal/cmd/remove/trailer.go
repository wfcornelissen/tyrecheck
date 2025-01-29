/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package remove

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wfcornelissen/tyrecheck/internal/entries"
)

// trailerCmd represents the trailer command
var trailerCmd = &cobra.Command{
	Use:   "trailer",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("trailer called")
		fleetNum := entries.ReadString("Enter fleet number: ")
		err := entries.RemoveTrailer(fleetNum)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	RemoveCmd.AddCommand(trailerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trailerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// trailerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
