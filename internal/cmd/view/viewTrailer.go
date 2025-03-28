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
		entries.ViewTrailer()
	},
}

func init() {
	ViewCmd.AddCommand(trailerCmd)

}
