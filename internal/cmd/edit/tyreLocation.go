package edit

import (
	"github.com/spf13/cobra"
	"github.com/wfcornelissen/tyrecheck/internal/entries"
)

// tyreLocationCmd represents the tyreLocation command
var tyreLocationCmd = &cobra.Command{
	Use:   "tyreLocation",
	Short: "Calls the tyre location entry function",
	Long: `Used as a subcommand for edit.
	Calls the tyre location entry function which asks for the tyre location.`,
	Run: func(cmd *cobra.Command, args []string) {
		TyreID := entries.ReadString("Enter the tyre ID: ")
		entries.EditLocation(TyreID)
	},
}

func init() {
	EditCmd.AddCommand(tyreLocationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tyreLocationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tyreLocationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
