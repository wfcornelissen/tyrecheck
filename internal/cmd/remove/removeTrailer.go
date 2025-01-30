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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trailerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// trailerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
