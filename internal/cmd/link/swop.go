/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package link

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wfcornelissen/tyrecheck/internal/entries"
)

// swopCmd represents the swop command
var SwopCmd = &cobra.Command{
	Use:   "swop",
	Short: "Swop trailers between two trucks",
	Long: `Used to assign trailers over between two trucks
	
	eg truck1, trailer1 and truck2, trailer2
	becomes truck1, trailer2 and truck2, trailer1`,
	Run: func(cmd *cobra.Command, args []string) {

		truckFleetNum1 := entries.ReadString("Truck Fleet Number 1: ")
		combo1, err := entries.CheckTruckTrailerCombo(truckFleetNum1)
		if err != nil {
			fmt.Println(err)
		}

		truckFleetNum2 := entries.ReadString("Truck Fleet Number 2: ")
		combo2, err := entries.CheckTruckTrailerCombo(truckFleetNum2)
		if err != nil {
			fmt.Println(err)
		}

		err = entries.SwopTruckTrailer(combo1.TruckFleetNum, combo2.TruckFleetNum, combo1.TrailerFleetNum, combo2.TrailerFleetNum)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	LinkCmd.AddCommand(SwopCmd)

}
