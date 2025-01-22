/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package add

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wfcornelissen/tyrecheck/internal/entries"
)

// trailerCmd represents the trailer command
var addTrailerCmd = &cobra.Command{
	Use:   "trailer",
	Short: "Calls the trailer entry function",
	Long: `Is used as subcommand for add to call
	the trailer entry function.
	
	It will ask for the following information:
	Fleet Number
	Vin
	Registration Number
	Make
	Model
	Year`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("trailer called")
		_, err := entries.AddTrailer()
		if err != nil {
			fmt.Println("Error adding trailer:", err)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trailerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// trailerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
