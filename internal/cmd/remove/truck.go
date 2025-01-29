/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package remove

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wfcornelissen/tyrecheck/internal/entries"
)

// truckCmd represents the truck command
var truckCmd = &cobra.Command{
	Use:   "truck",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("truck called")
		fleetNum := entries.ReadString("Enter fleet number: ")
		err := entries.RemoveTruck(fleetNum)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	RemoveCmd.AddCommand(truckCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// truckCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// truckCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
