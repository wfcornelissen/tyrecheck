package view

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wfcornelissen/tyrecheck/internal/entries"
)

// tyreCmd represents the tyre command
var tyreCmd = &cobra.Command{
	Use:   "tyre",
	Short: "View tyre details",
	Long:  `Displays details of tyre requested according to tyreID.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tyre called")
		entries.ViewTyre()
	},
}

func init() {
	ViewCmd.AddCommand(tyreCmd)

}
