package work

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wfcornelissen/tyrecheck/internal/entries"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Check: Calling check tyre")
		fleetNum := entries.ReadString("Please enter fleet number: ")
		if fleetNum == "" {
			fmt.Println("Error: Fleet number cannot be empty")
			return
		}
		tyrePosition := entries.ReadString("Please enter tyre Position: ")
		if tyrePosition == "" {
			fmt.Println("Error: Tyre position cannot be empty")
			return
		}
		tyrePosition = fleetNum + tyrePosition

		err := entries.CheckTyre(tyrePosition)
		if err != nil {
			fmt.Println("Error checking tyre:", err)
		}
	},
}

func init() {
	WorkCmd.AddCommand(checkCmd)

}
