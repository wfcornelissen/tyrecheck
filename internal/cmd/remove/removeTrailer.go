package remove

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wfcornelissen/tyrecheck/internal/entries"
)

// Finished
// trailerCmd represents the trailer command
var trailerCmd = &cobra.Command{
	Use:   "trailer",
	Short: "Remove a trailer from the database",
	Long:  `Sets a trailer scrap status to true based on fleet number input given by user`,
	Run: func(cmd *cobra.Command, args []string) {
		fleetNum := entries.ReadString("Enter fleet number: ")
		err := entries.RemoveTrailer(fleetNum)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	RemoveCmd.AddCommand(trailerCmd)

}
