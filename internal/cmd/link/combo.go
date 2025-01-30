/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package link

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wfcornelissen/tyrecheck/internal/entries"
)

// comboCmd represents the combo command
var comboCmd = &cobra.Command{
	Use:   "combo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("combo called")
		truckFleetNum := entries.ReadString("Truck Fleet Number: ")
		trailerFleetNum := entries.ReadString("Trailer Fleet Number: ")
		err := entries.LinkTruckTrailer(truckFleetNum, trailerFleetNum)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	LinkCmd.AddCommand(comboCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// comboCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// comboCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
