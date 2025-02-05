package remove

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wfcornelissen/tyrecheck/internal/entries"
)

// Finished
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

}
