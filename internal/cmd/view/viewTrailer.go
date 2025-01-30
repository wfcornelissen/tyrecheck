package view

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wfcornelissen/tyrecheck/internal/entries"
)

// trailerCmd represents the trailer command
var trailerCmd = &cobra.Command{
	Use:   "trailer",
	Short: "View trailer details",
	Long:  `Displays details of trailer requested according to fleet number.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("trailer called")
		fleetNum := entries.ReadString("Enter fleet number: ")
		err := entries.ViewTrailer(fleetNum)
		if err != nil {
			fmt.Println("Error viewing trailer:", err)
		}
	},
}

func init() {
	ViewCmd.AddCommand(trailerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trailerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// trailerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
