package work

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wfcornelissen/tyrecheck/internal/entries"
)

// fitCmd represents the fit command
var fitCmd = &cobra.Command{
	Use:   "fit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fit called")
		truckFleetNum := entries.ReadString("Please enter truck fleet number: ")
		if truckFleetNum == "" {
			fmt.Println("Error: Truck fleet number cannot be empty")
			return
		}
		newOrOld := entries.ConfirmEntry("Is the tyre being newly supplied?")
		if newOrOld {
			tyre, err := entries.AddTyre()
			if err != nil {
				fmt.Println("Error fitting tyre:", err)
			}
			if entries.ConfirmEntry(tyre) {
				entries.AssignTyre(truckFleetNum, tyre.ID)
			}
		} else if !newOrOld {
			tyreID := entries.ReadString("Please enter tyre ID: ")
			if tyreID == "" {
				fmt.Println("Error: Tyre ID cannot be empty")
				return
			}
			_, err := entries.ViewTyre(tyreID)
			if err != nil {
				fmt.Println("Error viewing tyre:", err)
			}
			if entries.ConfirmEntry("Are you sure you want to fit this tyre?") {
				entries.AssignTyre(truckFleetNum, tyreID)
			} else {
				fmt.Println("Tyre not fitted")
			}
		} else {
			fmt.Println("Invalid input")
		}
	},
}

func init() {
	WorkCmd.AddCommand(fitCmd)

}
