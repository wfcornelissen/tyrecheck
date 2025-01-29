package add

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wfcornelissen/tyrecheck/internal/entries"
)

// truckCmd represents the truck command
var addTruckCmd = &cobra.Command{
	Use:   "truck",
	Short: "Calls the truck entry function",
	Long: `Is used as subcommand for add to call
	the truck entry function.
	
	It will ask for the following information:
	Fleet Number
	VIN
	Registration Number
	Make
	Model
	Year
	Odometer
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("truck called") //Debug
		_, err := entries.AddTruck()
		if err != nil {
			fmt.Println("Error adding truck:", err)
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// truckCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// truckCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
