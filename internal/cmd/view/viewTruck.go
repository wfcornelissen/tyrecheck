package view

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wfcornelissen/tyrecheck/internal/entries"
)

// truckCmd represents the truck command
var truckCmd = &cobra.Command{
	Use:   "truck",
	Short: "View truck details",
	Long:  `Displays details of truck requested according to fleet number.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("truck called")
		fleetNum := entries.ReadString("Enter fleet number: ")
		err := entries.ViewTruck(fleetNum)
		if err != nil {
			fmt.Println("Error viewing truck:", err)
		}
	},
}

func init() {
	ViewCmd.AddCommand(truckCmd)

}
